package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type Config struct {
	DB  Database
	Jwt JWT
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
	// читаем переменные окружения из фпйла
	// если не прочитали возвращаем ошибку

	if err := godotenv.Load("./.env.example"); err != nil {
		return nil, fmt.Errorf("failed to load .env: %w", err)
	}

	var cfg Config

	// парсим переменные в структуру конфига
	// если не смогли распарсить возвращаем ошибку
	err := cleanenv.ReadEnv(&cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to read env: %w", err)
	}

	return &cfg, nil
}
