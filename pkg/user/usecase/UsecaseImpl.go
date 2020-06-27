package usecase

import (
	"context"

	"github.com/dhuki/Rest-Api-Golang/common"
	"github.com/dhuki/Rest-Api-Golang/pkg/user/domain/entity"
	"github.com/dhuki/Rest-Api-Golang/pkg/user/domain/repo"
	"github.com/dhuki/Rest-Api-Golang/pkg/user/domain/service"
	"github.com/dhuki/Rest-Api-Golang/pkg/user/presenter/model"
	"github.com/dhuki/Rest-Api-Golang/validation"
	"github.com/google/uuid"
)

type usecaseImpl struct {
	userRepo    repo.UserRepo
	userService service.UserService
}

func NewUsecase(userRepo repo.UserRepo, userService service.UserService) Usecase {
	return usecaseImpl{
		userRepo:    userRepo,
		userService: userService,
	}
}

func (u usecaseImpl) CreateUserUsecase(ctx context.Context, req model.CreateUserRequest) (common.BaseResponse, error) {
	var response common.BaseResponse
	{
		user := entity.User{
			ID:       uuid.New().String(),
			Username: req.Username,
			Password: req.Password,
			Role:     req.Role,
		}

		err := u.userRepo.CreateUser(ctx, user)
		if err != nil {
			return common.BaseResponse{}, err
		}

		token, err := validation.GenerateToken(user.ID, user.Role)
		if err != nil {
			return common.BaseResponse{}, err
		}

		userResponse := model.CreateUserResponse{
			Token:    token,
			Username: user.Username,
			Role:     user.Role,
		}

		response.Data = userResponse
	}
	return response, nil
}
