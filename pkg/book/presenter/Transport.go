package presenter

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dhuki/Rest-Api-Golang/common"
	"github.com/dhuki/Rest-Api-Golang/pkg/book/presenter/model"
	"github.com/dhuki/Rest-Api-Golang/pkg/book/usecase"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/go-kit/kit/transport"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func NewHttpServer(usecase usecase.Usecase, logger log.Logger) http.Handler {
	r := mux.NewRouter()
	r.Use(commonMiddleware)

	// setting up error log for internal and FE
	// ServerOption is func() type of data
	option := []httptransport.ServerOption{
		httptransport.ServerErrorEncoder(errorEncoder),
		// httptransport.ServerErrorHandler(transport.NewLogErrorHandler(level.Error(logger))), // default error internal logger
		httptransport.ServerErrorHandler(errorHandlerCustom(logger)),
	}

	r.Handle("/demo/api/books", httptransport.NewServer(
		MakeCreateBookEndpoint(usecase),
		model.DecodeCreateBookRequest,
		model.EncodeResponse,
		option...,
	)).Methods("POST")

	// r.Methods("GET").Path("").Handler(httptransport.NewServer())
	// r.Methods("GET").Path("/{id}").Handler(httptransport.NewServer())
	// r.Methods("POST").Path("/").Handler(httptransport.NewServer(
	// 	MakeCreateBookEndpoint(usecase),
	// 	model.DecodeCreateBookRequest,
	// 	model.EncodeResponse,
	// 	option...,
	// ))
	// r.Methods("PUT").Path("").Handler(httptransport.NewServer())
	// r.Methods("DELETE").Path("").Handler(httptransport.NewServer())

	return r
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json") // set up content-type of json at header for response
		fmt.Println(r.Method)
		fmt.Println(r.URL)
		next.ServeHTTP(w, r)
	})
}

// error for client
func errorEncoder(_ context.Context, err error, w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(common.BaseResponse{
		Message: err.Error(),
	})
}

// error custom for internal logger
func errorHandlerCustom(logger log.Logger) transport.ErrorHandlerFunc {
	return func(ctx context.Context, err error) {
		level.Error(logger).Log(
			"message", err,
			"description", "Transport error occured",
			"solution", "check usecase method, and dependency library")
	}
}
