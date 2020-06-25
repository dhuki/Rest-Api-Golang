package infrastructure

import (
	"context"

	"github.com/dhuki/Rest-Api-Golang/pkg/user/domain/entity"
	"github.com/dhuki/Rest-Api-Golang/pkg/user/domain/repo"
	"github.com/jinzhu/gorm"
)

type userRepoImpl struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) repo.UserRepo {
	return userRepoImpl{
		db: db,
	}
}

func (u userRepoImpl) CreateUser(ctx context.Context, user entity.User) error {
	db := u.db.Create(&user)
	if db.Error != nil {
		return db.Error
	}
	return nil
}
