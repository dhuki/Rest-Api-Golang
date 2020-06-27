package repo

import (
	"context"

	"github.com/dhuki/Rest-Api-Golang/pkg/book/domain/entity"
)

type BookRepo interface {
	Create(ctx context.Context, book entity.Book) error
	Update(ctx context.Context, book entity.Book) error
	GetBook(ctx context.Context, id string) (entity.Book, error)
	GetBooks(ctx context.Context) ([]entity.Book, error)
	DeleteBook(ctx context.Context, id string) error
}
