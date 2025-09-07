package server

import (
	"log"

	"github.com/pitipon/shop-microservices/modules/player/playerHandler"
	"github.com/pitipon/shop-microservices/modules/player/playerRepository"
	"github.com/pitipon/shop-microservices/modules/player/playerUsecase"

	playerPb "github.com/pitipon/shop-microservices/modules/player/playerPb"
	grpccon "github.com/pitipon/shop-microservices/pkg/gprccon"
)

func (s *server) playerService() {
	repo := playerRepository.NewPlayerRepository(s.db)
	usecase := playerUsecase.NewPlayerUsecase(repo)
	httpHandler := playerHandler.NewPlayerHandler(s.cfg, usecase)
	grpcHandler := playerHandler.NewPlayerGrpcHandler(usecase)
	queueHandler := playerHandler.NewPlayerQueueHandler(usecase)

	// start gRPC server
	go func() {
		grpcServer, lis := grpccon.NewGrpcServer(&s.cfg.Jwt, s.cfg.Grpc.PlayerUrl)

		playerPb.RegisterPlayerGrpcServiceServer(grpcServer, grpcHandler)

		log.Printf("Player gPRC server listening on %s", s.cfg.Grpc.PlayerUrl)
		grpcServer.Serve(lis)
	}()

	_ = grpcHandler
	_ = queueHandler

	player := s.app.Group("/player/v1")

	// Health check
	player.GET("", s.healthCheckService)

	player.POST("/player/register", httpHandler.CreatePlayer)
	player.GET("/player/:player_id", httpHandler.FindOnePlayerProfile)
	player.POST("/player/add-money", httpHandler.AddPlayerMoney)
	player.GET("/player/account/:player_id", httpHandler.GetPlayerSavingAccount)
}
