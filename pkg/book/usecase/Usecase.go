package usecase

import (
	"context"

	"github.com/dhuki/Rest-Api-Golang/common"
	"github.com/dhuki/Rest-Api-Golang/pkg/book/presenter/model"
)

type Usecase interface {
	CreateBookUsecase(context.Context, model.CreateBookRequest) (common.BaseResponse, error)
}
