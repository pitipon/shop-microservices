package server

import (
	"github.com/pitipon/shop-microservices/modules/item/itemHandler"
	"github.com/pitipon/shop-microservices/modules/item/itemRepository"
	"github.com/pitipon/shop-microservices/modules/item/itemUsecase"
)

func (s *server) itemService() {
	repo := itemRepository.NewItemRepository(s.db)
	usecase := itemUsecase.NewItemUsecase(repo)
	httpHandler := itemHandler.NewItemHttpHandler(s.cfg, usecase)
	grpcHandler := itemHandler.NewItemGrpcHandler(usecase)

	_ = httpHandler
	_ = grpcHandler

	item := s.app.Group("/item_v1")

	// Health check
	item.GET("/health", s.healthCheckService)
}
