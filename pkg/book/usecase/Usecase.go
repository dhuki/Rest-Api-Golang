package usecase

import (
	"context"

	"github.com/dhuki/Rest-Api-Golang/common"
	"github.com/dhuki/Rest-Api-Golang/pkg/book/domain/entity"
)

type Usecase interface {
	CreateBookUsecase(context.Context, entity.Book) (common.BaseResponse, error)
	UpdateBookUsecase(context.Context, entity.Book) (common.BaseResponse, error)
	GetBookUsecase(context.Context, entity.Book) (common.BaseResponse, error)
	GetBooksUsecase(context.Context) (common.BaseResponse, error)
	DeleteBookUsecase(context.Context, entity.Book) (common.BaseResponse, error)
}
