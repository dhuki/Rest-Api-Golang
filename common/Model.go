package common

import (
	"context"
	"net/http"
)

// BaseResponse is type of struct that wrap response data
type BaseResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Code    int         `json:"code"`
}

// ListRequest is common func to encode get list request
func ListRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	return nil, nil
}
