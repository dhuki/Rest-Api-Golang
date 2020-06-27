package presenter

import (
	"context"

	"github.com/dhuki/Rest-Api-Golang/pkg/book/domain/entity"
	"github.com/dhuki/Rest-Api-Golang/pkg/book/usecase"
	"github.com/go-kit/kit/endpoint"
)

func MakeCreateBookEndpoint(usecase usecase.Usecase) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(entity.Book) // assertion to change type
		response, err := usecase.CreateBookUsecase(ctx, req)
		return response, err
	}
}

func MakeUpdateBookEndpoint(usecase usecase.Usecase) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(entity.Book) // assertion to change type
		response, err := usecase.UpdateBookUsecase(ctx, req)
		return response, err
	}
}

func MakeGetBookEndpoint(usecase usecase.Usecase) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(entity.Book)
		response, err := usecase.GetBookUsecase(ctx, req)
		return response, err
	}
}

func MakeGetBooksEndpoint(usecase usecase.Usecase) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		response, err := usecase.GetBooksUsecase(ctx)
		return response, err
	}
}

func MakeDeleteBookEndpoint(usecase usecase.Usecase) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(entity.Book)
		response, err := usecase.DeleteBookUsecase(ctx, req)
		return response, err
	}
}
