package authUsecase

import "github.com/pitipon/shop-microservices/modules/auth/authRepository"

type (
	AuthUsercaseService interface{}

	authUsecase struct {
		authRepository authRepository.AuthRepositoryService
	}
)

func NewAuthUsercase(authRepository authRepository.AuthRepositoryService) AuthUsercaseService {
	return &authUsecase{authRepository}
}
