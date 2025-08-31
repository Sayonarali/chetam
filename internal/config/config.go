package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"

	"github.com/joho/godotenv"
)

type Config struct {
	DB     Database
	Jwt    JWT
	Server Server
}

type Server struct {
	Port string `env:"SRV_PORT"`
}

type Database struct {
	Host     string `env:"DB_HOST"`
	Port     string `env:"DB_ORT"`
	User     string `env:"DB_USER"`
	Password string `env:"DB_PASSWORD"`
	Name     string `env:"DB_NAME"`
}

type JWT struct {
	SecretKey string `env:"JWT_SECRET"`
	Sms       string `env:"JWT_SMS"`
	Phone     string `env:"JWT_PHONE"`
}

func Load() (*Config, error) {
	if err := godotenv.Load("./.env"); err != nil {
		return nil, fmt.Errorf("failed to load .env: %w", err)
	}

	var cfg Config

	err := cleanenv.ReadEnv(&cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to read env: %w", err)
	}

	return &cfg, nil
}
