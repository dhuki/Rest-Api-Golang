package service

import (
	"context"

	"github.com/dhuki/Rest-Api-Golang/pkg/book/domain/entity"
	"github.com/dhuki/Rest-Api-Golang/pkg/book/domain/repo"
)

type bookServiceImpl struct {
	repo repo.BookRepo
}

func NewBookService(repo repo.BookRepo) BookService {
	return bookServiceImpl{
		repo: repo,
	}
}

func (b bookServiceImpl) Create(ctx context.Context, book entity.Book) error {
	return nil
}
