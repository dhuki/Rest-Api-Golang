package service

import (
	"context"

	"github.com/dhuki/Rest-Api-Golang/pkg/book/domain/entity"
)

type BookService interface {
	Create(context.Context, entity.Book) error
}
