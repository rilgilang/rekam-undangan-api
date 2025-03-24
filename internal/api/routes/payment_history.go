package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rilgilang/kosan-api/config/dotenv"
	"github.com/rilgilang/kosan-api/internal/api/handlers"
	"github.com/rilgilang/kosan-api/internal/service"
)

func PaymentHistoryRoutes(app fiber.Router, cfg *dotenv.Config, service service.PaymentHistoryService) {
	app.Get("/payments/:room_id", handlers.GetAllPaymentHistoryByRoomId(cfg, service))
	app.Post("/payments", handlers.CreatePaymentHistoryByRoomId(cfg, service))
}
