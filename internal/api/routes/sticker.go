package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rilgilang/sticker-collection-api/config/dotenv"
	"github.com/rilgilang/sticker-collection-api/internal/api/handlers"
	"github.com/rilgilang/sticker-collection-api/internal/middlewares/jwt"
	"github.com/rilgilang/sticker-collection-api/internal/service"
)

func StickerRoutes(app fiber.Router, cfg *dotenv.Config, authMiddleware jwt.AuthMiddleware, service service.StickerService) {
	app.Get("/sticker", handlers.Stickers(cfg, service))
}
