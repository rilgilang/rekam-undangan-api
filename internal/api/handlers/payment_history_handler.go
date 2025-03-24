package handlers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/rilgilang/kosan-api/config/dotenv"
	"github.com/rilgilang/kosan-api/internal/api/presenter"
	"github.com/rilgilang/kosan-api/internal/entities"
	"github.com/rilgilang/kosan-api/internal/pkg/logger"
	"github.com/rilgilang/kosan-api/internal/service"
	"net/http"
)

func GetAllPaymentHistoryByRoomId(cfg *dotenv.Config, service service.PaymentHistoryService) fiber.Handler {
	return func(c *fiber.Ctx) error {

		var (
			ctx = c.Context()
		)

		roomId := c.Params("room_id")

		serv := service.GetAllPaymentHistoriesByRoomId(ctx, roomId)
		if serv.Code != 200 {
			c.Status(serv.Code)
			return c.JSON(presenter.ErrorResponse(serv.Errors))
		}

		c.Status(200)
		return c.JSON(presenter.SuccessResponse(serv.Data))
	}
}

func CreatePaymentHistoryByRoomId(cfg *dotenv.Config, service service.PaymentHistoryService) fiber.Handler {
	return func(c *fiber.Ctx) error {

		var (
			ctx     = c.Context()
			log     = logger.NewLog("get_room_handler", cfg.LoggerEnable)
			payload = entities.CreatePaymentHistoryPayload{}
		)

		err := c.BodyParser(&payload)
		if err != nil {
			log.Error(fmt.Sprintf(`error parsing request body got %s`, err))
			c.Status(http.StatusUnprocessableEntity)
			return c.JSON(presenter.ErrorResponse(err))
		}

		serv := service.CreatePaymentHistory(ctx, payload)
		if serv.Code != 200 {
			c.Status(serv.Code)
			return c.JSON(presenter.ErrorResponse(serv.Errors))
		}

		c.Status(200)
		return c.JSON(presenter.SuccessResponse(serv.Data))
	}
}
