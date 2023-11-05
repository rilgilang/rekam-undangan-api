package handlers

import (
	"digital_sekuriti_indonesia/config/yaml"
	"digital_sekuriti_indonesia/internal/api/presenter"
	"digital_sekuriti_indonesia/internal/consts"
	"digital_sekuriti_indonesia/internal/helper"
	"digital_sekuriti_indonesia/internal/pkg/logger"
	"digital_sekuriti_indonesia/internal/service"
	"github.com/gofiber/fiber/v2"
)

func GetProfile(cfg *yaml.Config, service service.AuthService) fiber.Handler {
	return func(c *fiber.Ctx) error {

		var (
			ctx = c.Context()
			log = logger.NewLog("get_profile_handler", cfg.Logger.Enable)
		)

		//log.Info(fmt.Sprintf(`start service login for user %s`, requestBody.Email))
		serv := service.GetProfile(ctx, helper.InterfaceToString(c.Locals(consts.UserId)))
		if serv.Code != 200 {
			//log.Error(fmt.Sprintf(`error on service login got %s`, serv.Errors))
			c.Status(serv.Code)
			return c.JSON(presenter.ErrorResponse(serv.Errors))
		}

		log.Info("get profile success")

		c.Status(200)
		return c.JSON(presenter.SuccessResponse(serv.Data))
	}
}
