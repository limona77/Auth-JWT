package app

import (
	"auth/config"
	"auth/internal/slogger"
	"github.com/gookit/slog"
	"log"
)

func Run(configPath string) {
	cfg, err := config.NewConfig(configPath)
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}
	slogger.SetLogger()
	slog.Info("cfg:", cfg)
}
