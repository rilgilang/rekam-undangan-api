package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/rilgilang/rekam-undangan-api/config/dotenv"
	"github.com/rilgilang/rekam-undangan-api/internal/api"
	"time"

	"log"
)

func main() {

	//cfg, err := yaml.NewConfig()
	//if err != nil {NewLoadConfig
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
		AllowMethods: "GET,POST",
	}), limiter.New(limiter.Config{
		Max:               100,
		Expiration:        30 * time.Second,
		LimiterMiddleware: limiter.SlidingWindow{},
	}))

	app = api.NewRouter(app, cfg)

	log.Fatal(app.Listen(fmt.Sprintf(`:%v`, cfg.AppPort)))
}
