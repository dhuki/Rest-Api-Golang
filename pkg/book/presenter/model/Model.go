package model

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// EncodeResponse from type struct to json
func EncodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	fmt.Println("woiit")
	return json.NewEncoder(w).Encode(response) // encode response type of struct to json
}

type (
	// CreateBookRequest is type struct for request
	CreateBookRequest struct {
		Title  string `json:"title"`
		Author string `json:"author"`
		Year   string `json:"year"`
	}
)

// DecodeCreateBookRequest from json to type struct
func DecodeCreateBookRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var request CreateBookRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		return nil, err
	}
	return request, nil
}
