package presenter

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/dhuki/Rest-Api-Golang/common"
	"github.com/dhuki/Rest-Api-Golang/pkg/book/presenter/model"
	"github.com/dhuki/Rest-Api-Golang/pkg/book/usecase"
	"github.com/dhuki/Rest-Api-Golang/validation"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func NewHttpServer(usecase usecase.Usecase, logger log.Logger) http.Handler {
	r := mux.NewRouter().StrictSlash(true)
	// r.Use(commonMiddleware)
	router := r.PathPrefix("/api").Subrouter()
	router.Use(commonMiddleware)
	router.Use(validationMiddleware(logger))

	// router2 := r.PathPrefix("/api").Subrouter()
	// router.Use(commonMiddleware)

	// setting up error log for internal and FE
	// ServerOption is func() type of data
	option := []httptransport.ServerOption{
		httptransport.ServerErrorEncoder(common.ErrorEncoder),
		// httptransport.ServerErrorHandler(transport.NewLogErrorHandler(level.Error(logger))), // default error internal logger
		httptransport.ServerErrorHandler(common.ErrorHandlerCustom(logger)),
	}

	// router.Methods("POST").Path("/books").Handler(httptransport.NewServer(
	// 	MakeCreateBookEndpoint(usecase),
	// 	model.DecodeCreateBookRequest,
	// 	model.EncodeResponse,
	// 	option...,
	// ))

	router.Methods("POST").Path("/").Handler(httptransport.NewServer(
		MakeCreateBookEndpoint(usecase),
		model.DecodeCreateBookRequest,
		model.EncodeResponse,
		option...,
	))

	router.Methods("PUT").Path("/").Handler(httptransport.NewServer(
		MakeUpdateBookEndpoint(usecase),
		model.DecodeUpdateBookRequest,
		model.EncodeResponse,
		option...,
	))

	router.Methods("GET").Path("/{id}").Handler(httptransport.NewServer(
		MakeGetBookEndpoint(usecase),
		model.DecodeGetBookRequest,
		model.EncodeResponse,
		option...,
	))

	router.Methods("GET").Path("/").Handler(httptransport.NewServer(
		MakeGetBooksEndpoint(usecase),
		common.ListRequest,
		model.EncodeResponse,
		option...,
	))

	router.Methods("DELETE").Path("/{id}").Handler(httptransport.NewServer(
		MakeDeleteBookEndpoint(usecase),
		model.DecodeDeleteBookRequest,
		model.EncodeResponse,
		option...,
	))

	// r.Handle("/demo/api/books", validationMiddleware(httptransport.NewServer(
	// 	MakeCreateBookEndpoint(usecase),
	// 	model.DecodeCreateBookRequest,
	// 	model.EncodeResponse,
	// 	option...,
	// ), logger)).Methods("POST")

	// r.Methods("POST").PathPrefix("/demo/api/books").Handler(httptransport.NewServer(
	// 	MakeCreateBookEndpoint(usecase),
	// 	model.DecodeCreateBookRequest,
	// 	model.EncodeResponse,
	// 	option...,
	// ))

	return router
}

func validationMiddleware(logger log.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()

			authHeader := r.Header.Get("Authorization")
			bearerToken := strings.Split(authHeader, " ") // bcs token return "Bearer xxxxxx"
			token := bearerToken[1]

			// fmt.Println(token)

			auth, err := validation.ParseToken(token)
			if err != nil {
				common.ErrorEncoder(ctx, err, w)
				level.Error(logger).Log("message", err, "description", "AUTHENTICATION & AUTHORIZATION")
				return
			}

			claim, ok := auth.Claims.(jwt.MapClaims) // assertion interface to map bcs it implement method interface
			if !ok {
				err = errors.New("Cannot make assertion")
				common.ErrorEncoder(ctx, err, w)
				level.Error(logger).Log("message", err, "description", "AUTHENTICATION & AUTHORIZATION")
				return
			}

			baseAuth := common.BaseAuth{
				ID:     claim["jti"].(string),
				URL:    r.URL.String(),
				Method: r.Method,
			}

			childCtx := context.WithValue(ctx, common.Auth, baseAuth) // making child of parent context with value inside it
			r = r.WithContext(childCtx)                               // bind child ctx with request

			next.ServeHTTP(w, r)
		})
	}
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json") // set up content-type of json at header for response
		next.ServeHTTP(w, r)
	})
}
