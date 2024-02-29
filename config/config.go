package config

import (
	"github.com/gookit/slog"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
	"path"
)

type (
	Config struct {
		HTTP `yaml:"http"`
		PG
	}
	HTTP struct {
		Port string `yaml:"port"`
	}
	PG struct {
		URL string ` env:"PG_URL_LOCALHOST"`
	}
)

func NewConfig(configPath string) *Config {
	cfg := &Config{}
	err := godotenv.Load()
	if err != nil {
		slog.Fatal("can't load env %w", err)
	}
	err = cleanenv.ReadConfig(path.Join("./", configPath), cfg)
	if err != nil {
		slog.Fatal("error reading config file: %w", err)
	}
	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		slog.Fatal("error reading env %w", err)
	}
	return cfg
}
