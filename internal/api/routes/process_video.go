package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rilgilang/rekam-undangan-api/config/dotenv"
	"github.com/rilgilang/rekam-undangan-api/internal/api/handlers"
	"github.com/rilgilang/rekam-undangan-api/internal/service"
)

func ProcessVideoRoutes(app fiber.Router, cfg *dotenv.Config, service service.VideoService) {
	app.Post("/process-video", handlers.ProcessVideo(cfg, service))
}
