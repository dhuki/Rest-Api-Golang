package model

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/dhuki/Rest-Api-Golang/pkg/book/domain/entity"
	"github.com/gorilla/mux"
)

// EncodeResponse from type struct to json
func EncodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response) // encode response type of struct to json
}

// type (
// 	// CreateBookRequest is type struct for request
// 	BookRequest struct {
// 		ID     string `json:"id"`
// 		Title  string `json:"title"`
// 		Author string `json:"author"`
// 		Year   string `json:"year"`
// 	}
// )

// DecodeCreateBookRequest from json to type struct
func DecodeCreateBookRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var request entity.Book
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		return nil, err
	}
	return request, nil
}

func DecodeUpdateBookRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var request entity.Book
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		return nil, err
	}
	return request, nil
}

func DecodeGetBookRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, errors.New("id is empty")
	}

	request := entity.Book{
		ID: id,
	}
	return request, nil
}

func DecodeDeleteBookRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, errors.New("id is empty")
	}

	request := entity.Book{
		ID: id,
	}
	return request, nil
}
