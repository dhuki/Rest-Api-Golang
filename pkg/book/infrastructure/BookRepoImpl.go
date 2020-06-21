package infrastructure

import (
	"context"

	"github.com/dhuki/Rest-Api-Golang/pkg/book/domain/entity"
	"github.com/dhuki/Rest-Api-Golang/pkg/book/domain/repo"
	"github.com/go-kit/kit/log"
	"github.com/jinzhu/gorm"
)

type bookRepoImpl struct {
	db     *gorm.DB
	logger log.Logger
}

func NewBookRepo(db *gorm.DB, logger log.Logger) repo.BookRepo {
	return bookRepoImpl{
		db:     db,
		logger: logger,
	}
}

func (b bookRepoImpl) Create(ctx context.Context, book entity.Book) error {
	db := b.db.Create(&book)
	if db.Error != nil {
		return db.Error
	}
	return nil
}
