package usecase

import (
	"context"

	"github.com/dhuki/Rest-Api-Golang/common"
	"github.com/dhuki/Rest-Api-Golang/pkg/user/presenter/model"
)

type Usecase interface {
	CreateUserUsecase(context.Context, model.CreateUserRequest) (common.BaseResponse, error)
}