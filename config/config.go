package config

import (
	"github.com/gookit/slog"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type (
	Config struct {
		HTTP
		PG
		JWT
	}
	HTTP struct {
		Port string `env:"HTTP_PORT"`
	}
	PG struct {
		URL string ` env:"PG_URL_LOCALHOST"`
	}
	JWT struct {
		SecretKeyAccess  []byte `env-required:"true" env:"JWT_ACCESS_SECRET"`
		SecretKeyRefresh []byte `env-required:"true" env:"JWT_REFRESH_SECRET"`
	}
)

func NewConfig() *Config {
	cfg := &Config{}
	err := godotenv.Load()
	if err != nil {
		slog.Fatal("can't load env %w", err)
	}
	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		slog.Fatal("error reading env %w", err)
	}
	return cfg
}
