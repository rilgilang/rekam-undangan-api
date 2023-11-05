package routes

import (
	"digital_sekuriti_indonesia/config/yaml"
	"digital_sekuriti_indonesia/internal/api/handlers"
	"digital_sekuriti_indonesia/internal/middlewares/jwt"
	"digital_sekuriti_indonesia/internal/service"
	"github.com/gofiber/fiber/v2"
)

func AuthRouter(app fiber.Router, cfg *yaml.Config, authMiddleware jwt.AuthMiddleware, service service.AuthService) {
	app.Post("/login", handlers.Login(cfg, service))
	app.Get("/profile", authMiddleware.ValidateToken(), handlers.GetProfile(cfg, service))
}
