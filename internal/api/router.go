package api

import (
	"digital_sekuriti_indonesia/bootstrap"
	"digital_sekuriti_indonesia/config/yaml"
	"digital_sekuriti_indonesia/internal/api/routes"
	"digital_sekuriti_indonesia/internal/middlewares/jwt"
	"digital_sekuriti_indonesia/internal/repositories"
	"digital_sekuriti_indonesia/internal/service"
	"digital_sekuriti_indonesia/migrations"
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

	migrations.AutoMigration(db)

	//repositories
	var (
		userRepo = repositories.NewUserRepo(db)
	)

	//middlewares
	var (
		authMidleware = jwt.NewAuthMiddleware(userRepo, cfg)
	)

	//service
	var (
		authService = service.NewAuthService(authMidleware, userRepo)
	)

	//group
	api := router.Group("/api")

	router.Get("/", func(c *fiber.Ctx) error {
		return c.Send([]byte("uwow"))
	})

	routes.AuthRouter(api, authService)

	//routes.HealthRouter(api)

	return router
}
