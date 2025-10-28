package main

import (
	"log"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/dyxgou/notas/cmd/api/routes"
	"github.com/dyxgou/notas/pkg/config"
	"github.com/dyxgou/notas/pkg/repositories/sqlite"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/healthcheck"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

const idleTimeout = 5 * time.Second

func main() {
	app := fiber.New(fiber.Config{
		IdleTimeout:  idleTimeout,
		ReadTimeout:  time.Second,
		WriteTimeout: time.Second,
	})

	// middlewares
	app.Use(helmet.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:4321,https://notas-frontend-theta.vercel.app/",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))
	app.Use(logger.New())
	app.Use(healthcheck.New())

	// connecting to database
	db := sqlite.ConnectClient(config.GetEnv("DB_PATH"))
	defer db.Close()

	slog.Info("database connected successfully")

	api := app.Group("/api")

	val := validator.New(validator.WithRequiredStructEnabled())

	// registering routes
	r := routes.NewRouter(db, val)

	r.RegisterUserGroup(api.Group("/student"))
	r.RegisterSubjectGroup(api.Group("/subject"))
	r.RegisterGradeGroup(api.Group("/grade"))
	r.RegisterNoteGroup(api.Group("/note"))
	r.RegisterReportGroup(api.Group("/report"))

	// Hooks
	hooks := app.Hooks()
	hooks.OnShutdown(func() error {
		slog.Info("database connection closed")
		return db.Close()
	})

	// initializing server
	go func() {
		if err := app.Listen(config.GetEnv("PORT")); err != nil {
			log.Fatal(err)
		}
	}()

	// Graceful shutdown
	c := make(chan os.Signal, 1)

	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	_ = <-c
	if err := app.Shutdown(); err != nil {
		log.Fatal(err)
	}
}
