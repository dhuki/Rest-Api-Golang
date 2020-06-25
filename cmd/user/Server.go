package user

import (
	"net/http"

	"github.com/dhuki/Rest-Api-Golang/pkg/user/domain/service"
	"github.com/dhuki/Rest-Api-Golang/pkg/user/infrastructure"
	"github.com/dhuki/Rest-Api-Golang/pkg/user/presenter"
	"github.com/dhuki/Rest-Api-Golang/pkg/user/usecase"
	"github.com/go-kit/kit/log"
	"github.com/jinzhu/gorm"
)

type server interface {
	Start() http.Handler
}

type userServer struct {
	db     *gorm.DB
	logger log.Logger
}

func NewServer(db *gorm.DB, logger log.Logger) server {
	return userServer{
		db:     db,
		logger: logger,
	}
}

func (u userServer) Start() http.Handler {
	var srv usecase.Usecase
	{
		infrastructure := infrastructure.NewUserRepo(u.db)
		service := service.NewUserService(infrastructure)
		srv = usecase.NewUsecase(infrastructure, service)
	}

	handler := presenter.NewServer(srv, u.logger)
	return handler
}
