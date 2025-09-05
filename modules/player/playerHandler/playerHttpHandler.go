package playerHandler

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/pitipon/shop-microservices/config"
	"github.com/pitipon/shop-microservices/modules/player"
	"github.com/pitipon/shop-microservices/modules/player/playerUsecase"
	"github.com/pitipon/shop-microservices/pkg/request"
	"github.com/pitipon/shop-microservices/pkg/response"
)

type (
	PlayerHttpHandlerService interface {
		CreatePlayer(c echo.Context) error
	}

	playerHttpHandler struct {
		cfg           *config.Config
		playerUsecase playerUsecase.PlayerUsecaseService
	}
)

func NewPlayerHandler(cfg *config.Config, playerUsecase playerUsecase.PlayerUsecaseService) PlayerHttpHandlerService {
	return &playerHttpHandler{
		cfg:           cfg,
		playerUsecase: playerUsecase,
	}
}

func (h *playerHttpHandler) CreatePlayer(c echo.Context) error {
	ctx := context.Background()

	wrapper := request.ContextWrapper(c)

	req := new(player.CreatePlayerReq)

	if err := wrapper.Bind(req); err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	res, err := h.playerUsecase.CreatePlayer(ctx, req)
	if err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	return response.SuccessResponse(c, http.StatusBadRequest, res)
}
