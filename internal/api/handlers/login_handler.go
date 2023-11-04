package handlers

import (
	"digital_sekuriti_indonesia/internal/api/presenter"
	"digital_sekuriti_indonesia/internal/entities"
	"digital_sekuriti_indonesia/internal/pkg/logger"
	"digital_sekuriti_indonesia/internal/service"
	"fmt"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func Login(service service.AuthService) fiber.Handler {
	return func(c *fiber.Ctx) error {

		var (
			requestBody = entities.User{}
			ctx         = c.Context()
			log         = logger.NewLog("login_handler")
		)

		err := c.BodyParser(&requestBody)
		if err != nil {
			log.Error(fmt.Sprintf(`error parsing request body got %s`, err))
			c.Status(http.StatusUnprocessableEntity)
			return c.JSON(presenter.ErrorResponse(err))
		}

		log.Info("validating request body")

		//validation using ozoo
		err = validation.ValidateStruct(&requestBody,
			validation.Field(&requestBody.Email, validation.Required),
			validation.Field(&requestBody.Password, validation.Required),
		)

		if err != nil {
			log.Warn(fmt.Sprintf(`error validating request body got %s`, err))
			c.Status(http.StatusBadRequest)
			return c.JSON(presenter.ErrorResponse(err))
		}

		log.Info(fmt.Sprintf(`start service login for user %s`, requestBody.Email))
		serv := service.Login(ctx, &requestBody)
		if serv.Code != 200 {
			log.Error(fmt.Sprintf(`error on service login got %s`, serv.Errors))
			c.Status(serv.Code)
			return c.JSON(presenter.ErrorResponse(serv.Errors))
		}

		log.Info("login success")

		c.Status(200)
		return c.JSON(presenter.SuccessResponse(serv.Data))
	}
}
