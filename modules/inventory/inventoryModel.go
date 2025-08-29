package inventory

import (
	"github.com/pitipon/shop-microservices/modules/item"
	"github.com/pitipon/shop-microservices/modules/models"
)

type (
	UpdateInventory struct {
		PlayerId string `json:"player_id" validate:"required,max=64"`
		ItemId   string `json:"item_id" validate:"required,max=64"`
	}

	ItemInInventory struct {
		InventoryId string `json:"inventory_id"`
		*item.ItemShowCase
	}

	PlayerInventory struct {
		PlayerId string `json:"player_id"`
		*models.PaginateRes
	}
)
