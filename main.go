package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/rilgilang/kosan-api/config/dotenv"
	"github.com/rilgilang/kosan-api/internal/api"

	"log"
)

func main() {

	//cfg, err := yaml.NewConfig()
	//if err != nil {
	//	log.Fatal(fmt.Sprintf(`read cfg yaml got error : %v`, err))
	//}

	cfg, err := dotenv.NewLoadConfig()
	if err != nil {
		log.Fatal(fmt.Sprintf(`read cfg .env got error : %v`, err))
	}

	app := fiber.New()

	// Apply middleware
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",
	}))

	app = api.NewRouter(app, cfg)

	log.Fatal(app.Listen(fmt.Sprintf(`:%v`, cfg.AppPort)))
}
