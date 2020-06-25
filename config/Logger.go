package config

import (
	"os"

	"github.com/dhuki/Rest-Api-Golang/custom"
	"github.com/go-kit/kit/log"
	logrus "github.com/sirupsen/logrus"
)

func NewLogger() log.Logger {
	// set up logrus
	var logLogrus *logrus.Logger
	{
		logLogrus = logrus.New()
		logLogrus.SetFormatter(&logrus.TextFormatter{
			TimestampFormat: "02-01-2006 15:04:05",
			FullTimestamp:   true,
			ForceColors:     true})
		logLogrus.SetOutput(os.Stdout)
	}

	// set up log associate with logrus
	var logger log.Logger
	{ // using closure, var closure to make more readable
		logger = custom.NewLogrusLogger(logLogrus)
		// encode to keyvals format
		// logger = log.NewLogfmtLogger(log.StdlibWriter{}) // os.Stderr return *File that implement Writer // or log.StdlibWriter{} using different format
		// it will blocked if there is multiple go routing use logger
		// only one goroutine will be allowed to log to the wrapped logger at a time.
		// until it available again
		logger = log.NewSyncLogger(logger)
		// it will add specs format error
		logger = log.With(logger,
			"caller", log.DefaultCaller,
			// "time: ", log.DefaultTimestamp,
			// "time: ", log.TimestampFormat(time.Now, common.FormatDate),
		)
	}
	return logger
}
