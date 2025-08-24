package server

import (
	"github.com/pitipon/shop-microservices/modules/player/playerHandler"
	"github.com/pitipon/shop-microservices/modules/player/playerRepository"
	"github.com/pitipon/shop-microservices/modules/player/playerUsecase"
)

func (s *server) playerService() {
	repo := playerRepository.NewPlayerRepository(s.db)
	usecase := playerUsecase.NewPlayerUsecase(repo)
	handler := playerHandler.NewPlayerHandler(s.cfg, usecase)
	grpcHandler := playerHandler.NewPlayerGrpcHandler(usecase)
	queueHandler := playerHandler.NewPlayerQueueHandler(usecase)

	_ = handler
	_ = grpcHandler
	_ = queueHandler

	player := s.app.Group("/player_v1")

	// Health check
	_ = player
}
