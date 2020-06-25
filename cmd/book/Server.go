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

type server interface {
	Start() http.Handler
}

type booksServer struct {
	db     *gorm.DB
	logger log.Logger
}

func NewServer(db *gorm.DB, logger log.Logger) server {
	return booksServer{
		db:     db,
		logger: logger,
	}
}

func (s booksServer) Start() http.Handler {
	var srv usecase.Usecase
	{
		infrastructure := infrastructure.NewBookRepo(s.db, s.logger)
		service := service.NewBookService(infrastructure)
		srv = usecase.NewUsecase(service, infrastructure, s.logger)
		srv = usecase.NewMiddleware(srv, s.logger)
	}

	handler := presenter.NewHttpServer(srv, s.logger)

	return handler
}
