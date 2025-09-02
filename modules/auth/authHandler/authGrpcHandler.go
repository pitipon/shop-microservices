package authHandler

import (
	"context"

	authPb "github.com/pitipon/shop-microservices/modules/auth/authPb"
	"github.com/pitipon/shop-microservices/modules/auth/authUsecase"
)

type (
	authGprcHandler struct {
		authPb.UnimplementedAuthGrpcServiceServer
		authUsecase authUsecase.AuthUsercaseService
	}
)

func NewAuthGrpcHandler(authUsecase authUsecase.AuthUsercaseService) *authGprcHandler {
	return &authGprcHandler{
		authUsecase: authUsecase,
	}
}

func (g *authGprcHandler) CredentialSearch(ctx context.Context, req *authPb.AccessTokenSearchReq) (*authPb.AccessTokenSearchRes, error) {
	return nil, nil
}

func (g *authGprcHandler) RolesCount(ctx context.Context, req *authPb.RolesCountReq) (*authPb.RolesCountRes, error) {
	return nil, nil
}
