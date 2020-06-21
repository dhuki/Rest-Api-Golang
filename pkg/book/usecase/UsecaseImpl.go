package usecase

import (
	"context"

	"github.com/dhuki/Rest-Api-Golang/common"
	"github.com/dhuki/Rest-Api-Golang/pkg/book/domain/entity"
	"github.com/dhuki/Rest-Api-Golang/pkg/book/domain/repo"
	"github.com/dhuki/Rest-Api-Golang/pkg/book/domain/service"
	"github.com/dhuki/Rest-Api-Golang/pkg/book/presenter/model"
	"github.com/go-kit/kit/log"
)

type usecaseImpl struct {
	bookService service.BookService
	bookRepo    repo.BookRepo
	logger      log.Logger
}

func NewUsecase(bookService service.BookService, bookRepo repo.BookRepo, logger log.Logger) Usecase {
	return usecaseImpl{
		bookService: bookService,
		bookRepo:    bookRepo,
		logger:      logger,
	}
}

func (u usecaseImpl) CreateBookUsecase(ctx context.Context, request model.CreateBookRequest) (common.BaseResponse, error) {
	var response common.BaseResponse
	{
		book := entity.Book{
			ID:     "b",
			Title:  request.Title,
			Author: request.Author,
			Year:   request.Year,
		}
		err := u.bookRepo.Create(ctx, book)
		if err != nil {
			return common.BaseResponse{}, err
		}
	}

	return response, nil
}
