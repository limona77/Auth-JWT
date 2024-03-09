package app

import (
	"auth/config"
	"auth/internal/controller"
	"auth/internal/repository"
	"auth/internal/service"
	"auth/internal/slogger"
	"auth/pkg/postgres"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gookit/slog"
)

const (
	ReadTimeout  = 3 * time.Second
	WriteTimeout = 3 * time.Second
)

func Run(configPath string) {
	slogger.SetLogger()

	slog.Info("init config")
	cfg := config.NewConfig(configPath)
	slog.Info("config ok")

	slog.Info("connecting to postgres")
	db := postgres.New(cfg.URL)
	defer db.Close()
	slog.Info("connect to postgres ok")

	slog.Info("init repositories")
	repositories := repository.NewRepositories(db)

	slog.Info("init services")
	deps := service.ServicesDeps{
		Repository:       repositories,
		SecretKeyAccess:  cfg.SecretKeyAccess,
		SecretKeyRefresh: cfg.SecretKeyRefresh,
	}

	services := service.NewServices(deps)

	fiberConfig := fiber.Config{
		ReadTimeout:  ReadTimeout,
		WriteTimeout: WriteTimeout,
	}
	app := fiber.New(fiberConfig)

	controller.NewRouter(app, services)

	slog.Info("starting fiber server")
	slog.Fatal(app.Listen(":" + cfg.Port))
}
