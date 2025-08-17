package authHandler

import "github.com/pitipon/shop-microservices/modules/auth/authUsecase"

type (
	authGprcHandler struct {
		authUsecase authUsecase.AuthUsercaseService
	}
)

func NewAuthGrpcHandler(authUsecase authUsecase.AuthUsercaseService) *authGprcHandler {
	return &authGprcHandler{
		authUsecase: authUsecase,
	}
}
