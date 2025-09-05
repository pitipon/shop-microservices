package server

import (
	"github.com/pitipon/shop-microservices/modules/payment/paymentHandler"
	"github.com/pitipon/shop-microservices/modules/payment/paymentRepository"
	"github.com/pitipon/shop-microservices/modules/payment/paymentUsecase"
)

func (s *server) paymentService() {
	repo := paymentRepository.NewPaymentRepository(s.db)
	usecase := paymentUsecase.NewPaymentUsecase(repo)
	handler := paymentHandler.NewPaymentHttpHandler(s.cfg, usecase)

	_ = handler

	payment := s.app.Group("/payment/v1")

	// Health Check
	payment.GET("", s.healthCheckService)

}
