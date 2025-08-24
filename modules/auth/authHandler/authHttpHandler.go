package authHandler

import (
	"github.com/pitipon/shop-microservices/config"
	"github.com/pitipon/shop-microservices/modules/auth/authUsecase"
)

type (
	AuthHttpHandlerService interface{}

	authHttpHandler struct {
		cfg         *config.Config
		authUsecase authUsecase.AuthUsercaseService
	}
)

func NewAuthHttpHandler(cfg *config.Config, authUsecase authUsecase.AuthUsercaseService) AuthHttpHandlerService {
	return &authHttpHandler{
		cfg:         cfg,
		authUsecase: authUsecase,
	}
}
