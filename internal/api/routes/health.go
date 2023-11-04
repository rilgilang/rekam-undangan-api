package routes

import (
	"github.com/gofiber/fiber/v2"
)

func HealthRouter(app fiber.Router) {
	app.Get("/health", handlers.Health())
}
