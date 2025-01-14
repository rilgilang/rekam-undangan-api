package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/rilgilang/sticker-collection-api/config/dotenv"
	"github.com/rilgilang/sticker-collection-api/internal/api"

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
	app.Use(cors.New())

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET, POST",
	}))

	app = api.NewRouter(cfg)

	log.Fatal(app.Listen(fmt.Sprintf(`:%v`, cfg.AppPort)))
}
