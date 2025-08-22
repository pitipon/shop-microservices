package playerHandler

import "github.com/pitipon/shop-microservices/modules/player/playerUsecase"

type (
	PlayerQueueHandlerService interface{}

	playerQueueHandler struct {
		playerUsecase playerUsecase.PlayerUsecaseService
	}
)

func NewPlayerQueueHandler(playerUsecase playerUsecase.PlayerUsecaseService) PlayerQueueHandlerService {
	return &playerQueueHandler{
		playerUsecase: playerUsecase,
	}
}
