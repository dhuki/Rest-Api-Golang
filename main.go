package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/dhuki/Rest-Api-Golang/cmd/book"
	"github.com/dhuki/Rest-Api-Golang/cmd/user"
	"github.com/dhuki/Rest-Api-Golang/common"
	"github.com/dhuki/Rest-Api-Golang/config"
	"github.com/go-kit/kit/log/level"
	"github.com/jinzhu/gorm"
)

func main() {

	// set up logger
	logger := config.NewLogger()
	level.Info(logger).Log("message", "service is running", "description", "SYSTEM IS RUNNING")
	defer level.Info(logger).Log("message", "service ended", "description", "SYSTEM IS SHUTDOWN")
	// end setup

	// set up database using gorm
	var db *gorm.DB
	{
		postgresServer := config.NewDatabase(logger)
		postgresDB, err := postgresServer.Start(common.DB_ENV_DIR)
		if err != nil {
			level.Error(logger).Log("message", err)
			os.Exit(-1)
		}
		db = postgresDB
	}
	defer db.Close() // close connection to db
	// end setup

	// initialize book server
	bookServer := book.NewServer(db, logger)

	// initialize book server
	userServer := user.NewServer(db, logger)

	errs := make(chan error)

	go func() {
		c := make(chan os.Signal)
		// SIGINT (Signal Interrupt (CTRL + C))
		// SIGTERM (Signal Terminated (KILL command))
		signal.Notify(c, syscall.SIGTERM, syscall.SIGINT) // insert to channel if there are centain signall
		errs <- fmt.Errorf("%s", <-c)
	}()

	go func() {
		mux := http.NewServeMux()
		// router
		// mux.Handle("/demo/api/books/", bookServer.Start())
		mux.Handle("/demo/api/users/", userServer.Start())
		mux.Handle("/demo/", http.StripPrefix("/demo", bookServer.Start()))

		errs <- http.ListenAndServe(":8080", mux) // return error when serve http
	}()

	level.Error(logger).Log("message", <-errs) // blocking until error occur
}
