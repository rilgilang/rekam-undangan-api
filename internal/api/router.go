package api

import (
	"digital_sekuriti_indonesia/bootstrap"
	"digital_sekuriti_indonesia/config/yaml"
	"digital_sekuriti_indonesia/internal/api/routes"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func NewRouter(cfg *yaml.Config) *fiber.App {
	router := fiber.New()

	db, err := bootstrap.DatabaseConnection(cfg)

	if err != nil {
		panic(fmt.Sprintf(`db connection error got : %v`, err))
	}

	fmt.Println("Database connection success!")

	//repositories

	//middlewares

	//service

	//group
	api := router.Group("/api")

	router.Get("/", func(c *fiber.Ctx) error {
		return c.Send([]byte("uwow"))
	})

	routes.HealthRouter(api)

	return router
}
