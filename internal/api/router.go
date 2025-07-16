package api

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/rilgilang/rekam-undangan-api/bootstrap"
	"github.com/rilgilang/rekam-undangan-api/config/dotenv"
	"github.com/rilgilang/rekam-undangan-api/internal/api/routes"
	"github.com/rilgilang/rekam-undangan-api/internal/pkg"
	"github.com/rilgilang/rekam-undangan-api/internal/repositories"
	"github.com/rilgilang/rekam-undangan-api/internal/service"
	"github.com/rilgilang/rekam-undangan-api/migrations"
)

func NewRouter(app *fiber.App, cfg *dotenv.Config) *fiber.App {

	db, err := bootstrap.DatabaseConnection(cfg)

	if err != nil {
		panic(fmt.Sprintf(`db connection error got : %v`, err))
	}

	//minioClient, err := bootstrap.NewMinio(cfg)

	// Redis client connect
	cache := pkg.NewCache(bootstrap.NewCache(cfg))

	fmt.Println("Database connection success!")

	migrations.AutoMigration(db)

	//repositories
	var (
		videoRepo = repositories.NewVideoRepo(db)
	)

	////middlewares
	//var (
	//	authMidleware = jwt.NewAuthMiddleware(userRepo, cfg)
	//)

	//service
	var (
		videoService = service.NewVideoService(videoRepo, cache, cfg)
	)

	//group
	api := app.Group("/api")

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Send([]byte("uwow"))
	})

	routes.ProcessVideoRoutes(api, cfg, videoService)

	//routes.HealthRouter(api)

	return app
}
