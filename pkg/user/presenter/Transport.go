package presenter

import (
	"net/http"

	"github.com/dhuki/Rest-Api-Golang/common"
	"github.com/dhuki/Rest-Api-Golang/pkg/user/presenter/model"
	"github.com/dhuki/Rest-Api-Golang/pkg/user/usecase"
	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func NewServer(usecase usecase.Usecase, logger log.Logger) http.Handler {
	r := mux.NewRouter()
	r.Use(commonMiddleware)

	option := []httptransport.ServerOption{
		httptransport.ServerErrorEncoder(common.ErrorEncoder),
		httptransport.ServerErrorHandler(common.ErrorHandlerCustom(logger)),
	}

	r.Handle("/demo/api/users/sign-up", httptransport.NewServer(
		MakeCreateUserRequest(usecase),
		model.DecodeCreateUserRequest,
		model.EncodeResponse,
		option...,
	)).Methods("POST")

	return r
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json") // set up content-type of json at header for response
		next.ServeHTTP(w, r)
	})
}
