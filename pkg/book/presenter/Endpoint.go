package presenter

import (
	"context"

	"github.com/dhuki/Rest-Api-Golang/pkg/book/presenter/model"
	"github.com/dhuki/Rest-Api-Golang/pkg/book/usecase"
	"github.com/go-kit/kit/endpoint"
)

func MakeCreateBookEndpoint(usecase usecase.Usecase) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(model.CreateBookRequest) // assertion to change type
		response, err := usecase.CreateBookUsecase(ctx, req)
		return response, err
	}
}
