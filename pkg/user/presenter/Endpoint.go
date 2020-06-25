package presenter

// this package is making connection between request from client and bussiness process
// return endpoint.Endpoint by gokit

import (
	"context"

	"github.com/dhuki/Rest-Api-Golang/pkg/user/presenter/model"
	"github.com/dhuki/Rest-Api-Golang/pkg/user/usecase"
	"github.com/go-kit/kit/endpoint"
)

func MakeCreateUserRequest(usecase usecase.Usecase) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(model.CreateUserRequest)
		response, err := usecase.CreateUserUsecase(ctx, req)
		if err != nil {
			return nil, err
		}
		return response, nil
	}
}
