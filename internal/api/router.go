package api

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/rilgilang/kosan-api/bootstrap"
	"github.com/rilgilang/kosan-api/config/dotenv"
	"github.com/rilgilang/kosan-api/internal/api/routes"
	"github.com/rilgilang/kosan-api/internal/repositories"
	"github.com/rilgilang/kosan-api/internal/service"
	"github.com/rilgilang/kosan-api/migrations"
)

func NewRouter(app *fiber.App, cfg *dotenv.Config) *fiber.App {

	db, err := bootstrap.DatabaseConnection(cfg)

	if err != nil {
		panic(fmt.Sprintf(`db connection error got : %v`, err))
	}

	//minioClient, err := bootstrap.NewMinio(cfg)

	fmt.Println("Database connection success!")

	migrations.AutoMigration(db)

	//repositories
	var (
		roomRepo           = repositories.NewRoomRepo(db)
		paymentHistoryRepo = repositories.NewPaymentHistoryRepo(db)
	)

	////middlewares
	//var (
	//	authMidleware = jwt.NewAuthMiddleware(userRepo, cfg)
	//)

	//service
	var (
		roomService           = service.NewRoomService(roomRepo, paymentHistoryRepo, cfg)
		paymentHistoryService = service.NewPaymentHistoryService(paymentHistoryRepo, roomRepo, cfg)
	)

	//group
	api := app.Group("/api")

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Send([]byte("uwow"))
	})

	routes.RoomRoutes(api, cfg, roomService)
	routes.PaymentHistoryRoutes(api, cfg, paymentHistoryService)

	//routes.HealthRouter(api)

	return app
}
