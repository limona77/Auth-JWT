package app

import (
	"auth/config"
	"auth/internal/slogger"
	"auth/pkg"
	"github.com/gookit/slog"
)

func Run(configPath string) {
	slogger.SetLogger()

	slog.Info("init config")
	cfg := config.NewConfig(configPath)
	slog.Info("config ok")

	slog.Info("connecting to postgres")
	pg := postgres.New(cfg.URL)
	defer pg.Pool.Close()
	slog.Info("connect to postgres ok")
}
