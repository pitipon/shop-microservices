package server

import (
	"github.com/pitipon/shop-microservices/modules/payment/paymentRepository"
	"github.com/pitipon/shop-microservices/modules/player/playerHandler"
	"github.com/pitipon/shop-microservices/modules/player/playerUsecase"
)

func (s *server) paymentService() {
	repo := paymentRepository.NewPaymentRepository(s.db)
	usecase := playerUsecase.NewPlayerUsecase(repo)
	handler := playerHandler.NewPlayerHandler(s.cfg, usecase)
	gprcHandler := playerHandler.NewPlayerGrpcHandler(usecase)
	queueHandler := playerHandler.NewPlayerQueueHandler(usecase)

	_ = handler
	_ = gprcHandler
	_ = queueHandler

	payment := s.app.Group("/payment_v1")

	// Health Check
	payment.GET("/health", s.healthCheckService)

}
