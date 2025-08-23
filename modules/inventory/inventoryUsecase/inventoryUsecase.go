package inventoryUsecase

import "github.com/pitipon/shop-microservices/modules/inventory/inventoryRepository"

type (
	InventoryUsecaseService interface{}

	inventoryUsecase struct {
		inventoryRepository inventoryRepository.InventoryRepositoryService
	}
)

func NewInventoryUsecase(inventoryRepository inventoryRepository.InventoryRepositoryService) InventoryUsecaseService {
	return &inventoryUsecase{
		inventoryRepository: inventoryRepository,
	}
}
