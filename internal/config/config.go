package config

import (
	"fmt"
	"os"
)

type Config struct {
	Port string
	DB   string
}

func Load() *Config {

	return &Config{
		Port: os.Getenv("APP_PORT"),
		DB: fmt.Sprintf(
			"postgres://%s:%s@%s:%s/%s?sslmode=disable",
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_HOST"),
			os.Getenv("DB_PORT"),
			os.Getenv("DB_NAME"),
		),
	}
}
