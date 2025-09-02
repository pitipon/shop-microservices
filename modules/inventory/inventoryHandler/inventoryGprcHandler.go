package inventoryHandler

import (
	"context"

	inventoryPb "github.com/pitipon/shop-microservices/modules/inventory/inventoryPb"
	"github.com/pitipon/shop-microservices/modules/inventory/inventoryUsecase"
)

type (
	inventoryGprcHandler struct {
		inventoryPb.UnimplementedInventoryGrpcServiceServer
		inventoryUsecase inventoryUsecase.InventoryUsecaseService
	}
)

func NewInventoryGrpcHandler(inventoryUsecase inventoryUsecase.InventoryUsecaseService) *inventoryGprcHandler {
	return &inventoryGprcHandler{
		inventoryUsecase: inventoryUsecase,
	}
}

func (g *inventoryGprcHandler) IsAvailableToSell(ctx context.Context, req *inventoryPb.IsAvailableToSellReq) (*inventoryPb.IsAvailableToSellRes, error) {
	return nil, nil
}
