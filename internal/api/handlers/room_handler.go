package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rilgilang/kosan-api/config/dotenv"
	"github.com/rilgilang/kosan-api/internal/api/presenter"
	"github.com/rilgilang/kosan-api/internal/pkg/logger"
	"github.com/rilgilang/kosan-api/internal/service"
)

func GetAllRoom(cfg *dotenv.Config, service service.RoomService) fiber.Handler {
	return func(c *fiber.Ctx) error {

		var (
			ctx = c.Context()
			log = logger.NewLog("get_room_handler", cfg.LoggerEnable)
		)

		//log.Info(fmt.Sprintf(`start service login for user %s`, requestBody.Email))
		serv := service.GetAllRooms(ctx)
		if serv.Code != 200 {
			//log.Error(fmt.Sprintf(`error on service login got %s`, serv.Errors))
			c.Status(serv.Code)
			return c.JSON(presenter.ErrorResponse(serv.Errors))
		}

		log.Info("get room success")

		c.Status(200)
		return c.JSON(presenter.SuccessResponse(serv.Data))
	}
}
