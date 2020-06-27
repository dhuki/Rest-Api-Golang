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

func (b bookRepoImpl) Update(ctx context.Context, book entity.Book) error {
	db := b.db.Save(&book)
	if db.Error != nil {
		return db.Error
	}
	return nil
}

func (b bookRepoImpl) GetBook(ctx context.Context, id string) (entity.Book, error) {
	book := entity.Book{
		ID: id,
	}

	db := b.db.First(&book)
	if db.Error != nil {
		return entity.Book{}, db.Error
	}
	return book, nil
}

func (b bookRepoImpl) GetBooks(ctx context.Context) ([]entity.Book, error) {
	var books []entity.Book
	db := b.db.Find(&books)
	if db.Error != nil {
		return []entity.Book{}, db.Error
	}
	return books, nil
}

func (b bookRepoImpl) DeleteBook(ctx context.Context, id string) error {
	book := entity.Book{
		ID: id,
	}

	db := b.db.Delete(&book)
	if db.Error != nil {
		return db.Error
	}
	return nil
}
