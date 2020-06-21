package book

import (
	"net/http"

	"github.com/dhuki/Rest-Api-Golang/pkg/book/domain/service"
	"github.com/dhuki/Rest-Api-Golang/pkg/book/infrastructure"
	"github.com/dhuki/Rest-Api-Golang/pkg/book/presenter"
	"github.com/dhuki/Rest-Api-Golang/pkg/book/usecase"
	"github.com/go-kit/kit/log"
	"github.com/jinzhu/gorm"
)

type server struct {
	db     *gorm.DB
	logger log.Logger
}

func NewServer(db *gorm.DB, logger log.Logger) server {
	return server{
		db:     db,
		logger: logger,
	}
}

func (s server) Start() http.Handler {
	var srv usecase.Usecase
	{
		infrastructure := infrastructure.NewBookRepo(s.db, s.logger)
		service := service.NewBookService(infrastructure)
		srv = usecase.NewUsecase(service, infrastructure, s.logger)
		srv = usecase.NewMiddleware(s.logger, srv)
	}

	handler := presenter.NewHttpServer(srv, s.logger)

	return handler
}
