package service

import "github.com/dhuki/Rest-Api-Golang/pkg/user/domain/repo"

type userServiceImpl struct {
	userRepo repo.UserRepo
}

func NewUserService(userRepo repo.UserRepo) UserService {
	return userServiceImpl{
		userRepo: userRepo,
	}
}
