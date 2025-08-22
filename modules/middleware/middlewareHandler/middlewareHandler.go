package middlewareHandler

import (
	"github.com/pitipon/shop-microservices/config"
	"github.com/pitipon/shop-microservices/modules/middleware/middlewareUsecase"
)

type (
	MiddlewareHandlerService interface{}

	middlewareHandler struct {
		cfg               *config.Config
		middlewareUsecase middlewareUsecase.MiddlewareUsecaseService
	}
)

func NewMiddlewareHandler(cfg *config.Config, middlewareUsecase middlewareUsecase.MiddlewareUsecaseService) MiddlewareHandlerService {
	return &middlewareHandler{
		cfg:               cfg,
		middlewareUsecase: middlewareUsecase,
	}
}
