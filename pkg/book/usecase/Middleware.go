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

func NewMiddleware(usecase Usecase, logger log.Logger) Usecase {
	return middleware{
		logger:  logger,
		usecase: usecase,
	}
}

func (m middleware) CreateBookUsecase(ctx context.Context, request model.CreateBookRequest) (response common.BaseResponse, err error) {
	defer func(begin time.Time) {
		// baseInfo := ctx.Value(common.Auth).(common.BaseAuth)
		level.Info(m.logger).Log(
			"description", "INTERCEPTOR",
			"took", time.Since(begin),
			// "url", baseInfo.URL,
			// "method", baseInfo.Method,
			"request", fmt.Sprintf("%+v", request), // givin output of struct to this -> attribute : value
			"response", fmt.Sprintf("%+v", response)) // givin output of struct to this -> attribute : value
	}(time.Now())
	return m.usecase.CreateBookUsecase(ctx, request)
}
