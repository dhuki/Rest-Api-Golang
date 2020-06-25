package common

import (
	"context"
	"net/http"
)

// BaseAuth is type of struct that hold user's information
type BaseAuth struct {
	ID     string
	URL    string
	Method string
}

// BaseResponse is type of struct that wrap response data
type BaseResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Code    int         `json:"code"`
}

func GenerateBaseResponse(success bool, message string, data interface{}, code int) BaseResponse {
	return BaseResponse{
		Success: success,
		Message: message,
		Data:    data,
		Code:    code,
	}
}

// ListRequest is common func to encode get list request
func ListRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	return nil, nil
}
