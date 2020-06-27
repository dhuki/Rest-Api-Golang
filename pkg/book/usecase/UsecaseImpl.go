package usecase

import (
	"context"

	"github.com/dhuki/Rest-Api-Golang/common"
	"github.com/dhuki/Rest-Api-Golang/pkg/book/domain/entity"
	"github.com/dhuki/Rest-Api-Golang/pkg/book/domain/repo"
	"github.com/dhuki/Rest-Api-Golang/pkg/book/domain/service"
	"github.com/go-kit/kit/log"
	"github.com/google/uuid"
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

func (u usecaseImpl) CreateBookUsecase(ctx context.Context, request entity.Book) (common.BaseResponse, error) {
	var response common.BaseResponse
	{
		book := entity.Book{
			ID:     uuid.New().String(),
			Title:  request.Title,
			Author: request.Author,
			Year:   request.Year,
		}
		err := u.bookRepo.Create(ctx, book)
		if err != nil {
			return common.BaseResponse{}, err
		}
		response.Success = true
		response.Message = common.Success
	}

	return response, nil
}

func (u usecaseImpl) UpdateBookUsecase(ctx context.Context, request entity.Book) (common.BaseResponse, error) {
	var response common.BaseResponse
	{
		book, err := u.bookRepo.GetBook(ctx, request.ID)
		if err != nil {
			return common.BaseResponse{}, err
		}

		book.Author = request.Author
		book.Title = request.Title
		book.Year = request.Year

		err = u.bookRepo.Update(ctx, book)
		if err != nil {
			return common.BaseResponse{}, err
		}

		response.Success = true
		response.Message = common.Success
	}
	return response, nil
}

func (u usecaseImpl) GetBookUsecase(ctx context.Context, request entity.Book) (common.BaseResponse, error) {
	var response common.BaseResponse
	{
		book, err := u.bookRepo.GetBook(ctx, request.ID)
		if err != nil {
			return common.BaseResponse{}, err
		}
		response.Data = book
		response.Success = true
		response.Message = common.Success
	}
	return response, nil
}

func (u usecaseImpl) GetBooksUsecase(ctx context.Context) (common.BaseResponse, error) {
	var response common.BaseResponse
	{
		books, err := u.bookRepo.GetBooks(ctx)
		if err != nil {
			return common.BaseResponse{}, nil
		}
		response.Data = books
		response.Success = true
		response.Message = common.Success
	}
	return response, nil
}

func (u usecaseImpl) DeleteBookUsecase(ctx context.Context, request entity.Book) (common.BaseResponse, error) {
	var response common.BaseResponse
	{
		err := u.bookRepo.DeleteBook(ctx, request.ID)
		if err != nil {
			return common.BaseResponse{}, err
		}
		response.Success = true
		response.Message = common.Success
	}
	return response, nil
}
