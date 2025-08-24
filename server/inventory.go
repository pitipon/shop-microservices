package server

import (
	"github.com/pitipon/shop-microservices/modules/inventory/inventoryHandler"
	"github.com/pitipon/shop-microservices/modules/inventory/inventoryRepository"
	"github.com/pitipon/shop-microservices/modules/inventory/inventoryUsecase"
)

func (s *server) inventoryService() {
	repo := inventoryRepository.NewInventoryRepository(s.db)
	usecase := inventoryUsecase.NewInventoryUsecase(repo)
	httpHandler := inventoryHandler.NewInventoryHttpHandler(s.cfg, usecase)
	queueHandler := inventoryHandler.NewInventoryQueueHandler(s.cfg, usecase)

	_ = httpHandler
	_ = queueHandler

	inventory := s.app.Group("/inventory_v1")

	// Health check
	_ = inventory
}
