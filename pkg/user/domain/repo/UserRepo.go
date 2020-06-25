package repo

import (
	"context"

	"github.com/dhuki/Rest-Api-Golang/pkg/user/domain/entity"
)

type UserRepo interface {
	CreateUser(ctx context.Context, user entity.User) error
}
