package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rilgilang/kosan-api/config/dotenv"
	"github.com/rilgilang/kosan-api/internal/api/handlers"
	"github.com/rilgilang/kosan-api/internal/service"
)

func RoomRoutes(app fiber.Router, cfg *dotenv.Config, service service.RoomService) {
	app.Get("/room", handlers.GetAllRoom(cfg, service))
}
