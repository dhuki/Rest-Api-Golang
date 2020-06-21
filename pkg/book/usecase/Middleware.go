package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/dhuki/Rest-Api-Golang/common"
	"github.com/dhuki/Rest-Api-Golang/pkg/book/presenter/model"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)

type middleware struct {
	logger  log.Logger
	usecase Usecase
}

func NewMiddleware(logger log.Logger, usecase Usecase) Usecase {
	return middleware{
		logger:  logger,
		usecase: usecase,
	}
}

func (m middleware) CreateBookUsecase(ctx context.Context, request model.CreateBookRequest) (response common.BaseResponse, err error) {
	defer func(begin time.Time) {
		level.Info(m.logger).Log(
			"description", "INTERCEPTOR",
			"took", time.Since(begin),
			"request", fmt.Sprintf("%+v", request),
			"response", fmt.Sprintf("%+v", response))
	}(time.Now())
	return m.usecase.CreateBookUsecase(ctx, request)
}
