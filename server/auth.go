package server

import (
	"log"

	"github.com/pitipon/shop-microservices/modules/auth/authHandler"

	authPb "github.com/pitipon/shop-microservices/modules/auth/authPb"
	"github.com/pitipon/shop-microservices/modules/auth/authRepository"
	"github.com/pitipon/shop-microservices/modules/auth/authUsecase"
	grpccon "github.com/pitipon/shop-microservices/pkg/gprccon"
)

func (s *server) authService() {
	repo := authRepository.NewAuthRepository(s.db)
	usecase := authUsecase.NewAuthUsercase(repo)
	httpHandler := authHandler.NewAuthHttpHandler(s.cfg, usecase)
	grpcHandler := authHandler.NewAuthGrpcHandler(usecase)

	// start gRPC server
	go func() {
		grpcServer, lis := grpccon.NewGrpcServer(&s.cfg.Jwt, s.cfg.Grpc.AuthUrl)

		authPb.RegisterAuthGrpcServiceServer(grpcServer, grpcHandler)

		log.Printf("Auth gPRC server listening on %s", s.cfg.Grpc.AuthUrl)
		grpcServer.Serve(lis)
	}()

	_ = httpHandler
	_ = grpcHandler

	auth := s.app.Group("/auth/v1")

	// Health Check
	auth.GET("", s.healthCheckService)

}
