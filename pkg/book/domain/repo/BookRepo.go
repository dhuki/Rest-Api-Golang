package repo

import (
	"context"

	"github.com/dhuki/Rest-Api-Golang/pkg/book/domain/entity"
)

type BookRepo interface {
	Create(ctx context.Context, book entity.Book) error
}
