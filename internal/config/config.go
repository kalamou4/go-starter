package config

import (
	"os"
)

type Config struct {
	ServerAddr  string
	DatabaseURL string
	JWTSecret   string
}

func Load() (*Config, error) {
	return &Config{
		ServerAddr:  os.Getenv("SERVER_ADDR"),
		DatabaseURL: os.Getenv("DATABASE_URL"),
		JWTSecret:   os.Getenv("JWT_SECRET"),
	}, nil
}
