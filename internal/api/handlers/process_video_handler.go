package handlers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/rilgilang/rekam-undangan-api/config/dotenv"
	"github.com/rilgilang/rekam-undangan-api/internal/api/presenter"
	"github.com/rilgilang/rekam-undangan-api/internal/entities"
	"github.com/rilgilang/rekam-undangan-api/internal/pkg/logger"
	"github.com/rilgilang/rekam-undangan-api/internal/service"
	"net/http"
)

//func GetAllProcessedVideo(cfg *dotenv.Config, service service.VideoService) fiber.Handler {
//	return func(c *fiber.Ctx) error {
//
//		var (
//			ctx = c.Context()
//		)
//
//		serv := service.GetAllPaymentHistoriesByRoomId(ctx, roomId)
//		if serv.Code != 200 {
//			c.Status(serv.Code)
//			return c.JSON(presenter.ErrorResponse(serv.Errors))
//		}
//
//		c.Status(200)
//		return c.JSON(presenter.SuccessResponse(serv.Data))
//	}
//}

func ProcessVideo(cfg *dotenv.Config, service service.VideoService) fiber.Handler {
	return func(c *fiber.Ctx) error {

		var (
			ctx     = c.Context()
			log     = logger.NewLog("process_video_handler", cfg.LoggerEnable)
			payload = entities.ProcessVideoPayload{}
		)

		err := c.BodyParser(&payload)
		if err != nil {
			log.Error(fmt.Sprintf(`error parsing request body got %s`, err))
			c.Status(http.StatusUnprocessableEntity)
			return c.JSON(presenter.ErrorResponse(err))
		}

		serv := service.ProcessVideo(ctx, payload.URL)
		if serv.Code != 200 {
			c.Status(serv.Code)
			return c.JSON(presenter.ErrorResponse(serv.Errors))
		}

		c.Status(200)
		return c.JSON(presenter.SuccessResponse(serv.Data))
	}
}
