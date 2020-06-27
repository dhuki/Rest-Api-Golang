package common

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-kit/kit/auth/jwt"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/go-kit/kit/transport"
	"github.com/jinzhu/gorm"
)

// ErrorEncoder is error handling for client
func ErrorEncoder(_ context.Context, err error, w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	var response BaseResponse
	{
		response.Success = false
		switch err {
		case gorm.ErrRecordNotFound:
			response.Message = err.Error()
		case jwt.ErrTokenInvalid:
			response.Message = "Error token invalid"
		default:
			response.Message = ErrInternalServerError
		}
		json.NewEncoder(w).Encode(response)
	}
}

// ErrorHandlerCustom is error handling for internal logger
func ErrorHandlerCustom(logger log.Logger) transport.ErrorHandlerFunc {
	return func(ctx context.Context, err error) {
		level.Error(logger).Log(
			"message", err,
			"description", "Transport error occured",
			"solution", "Please check encode/decode body, usecase method, and dependency library")
	}
}
