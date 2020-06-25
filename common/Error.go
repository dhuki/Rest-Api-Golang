package common

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/go-kit/kit/transport"
)

// ErrorEncoder is error handling for client
func ErrorEncoder(_ context.Context, err error, w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(BaseResponse{
		Message: err.Error(),
	})
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
