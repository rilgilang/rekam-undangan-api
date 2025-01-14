package api

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/rilgilang/sticker-collection-api/bootstrap"
	"github.com/rilgilang/sticker-collection-api/config/dotenv"
	"github.com/rilgilang/sticker-collection-api/internal/api/routes"
	"github.com/rilgilang/sticker-collection-api/internal/middlewares/jwt"
	"github.com/rilgilang/sticker-collection-api/internal/repositories"
	"github.com/rilgilang/sticker-collection-api/internal/service"
	"github.com/rilgilang/sticker-collection-api/migrations"
)

func NewRouter(cfg *dotenv.Config) *fiber.App {
	router := fiber.New()

	db, err := bootstrap.DatabaseConnection(cfg)

	if err != nil {
		panic(fmt.Sprintf(`db connection error got : %v`, err))
	}

	fmt.Println("Database connection success!")

	migrations.AutoMigration(db)

	//repositories
	var (
		userRepo    = repositories.NewUserRepo(db)
		stickerRepo = repositories.NewStickerRepo(db)
	)

	//middlewares
	var (
		authMidleware = jwt.NewAuthMiddleware(userRepo, cfg)
	)

	//service~
	var (
		authService    = service.NewAuthService(authMidleware, userRepo, cfg)
		stickerService = service.NewStickerService(stickerRepo, cfg)
	)

	//group
	api := router.Group("/api")

	router.Get("/", func(c *fiber.Ctx) error {
		return c.Send([]byte("uwow"))
	})

	routes.AuthRouter(api, cfg, authMidleware, authService)
	routes.StickerRoutes(api, cfg, authMidleware, stickerService)

	//routes.HealthRouter(api)

	return router
}
