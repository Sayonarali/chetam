package main

import (
	"fmt"
	"log/slog"
	"os"

	"chetam/internal/config"
	"chetam/internal/db/client/postgres"
	"chetam/internal/services/auth"
	"chetam/internal/transport/repository"
	"chetam/pkg/logger"
)

func main() {
	// инициализируем логгер
	lg := logger.New()
	// логгер по умолчанию
	slog.SetDefault(lg)

	// загружаем конфиг из internal/config
	cfg, err := config.Load()
	if err != nil {
		lg.Error(
			"failed to load config",
			slog.String("error", err.Error()),
		)
	}

	fmt.Printf("config: %+v\n", cfg)

	// инициализируем клиент для базы данных
	// передаем туда конфиг
	// если ошибка = останавливаем программу
	client, err := postgres.NewClient(cfg)
	if err != nil {
		lg.Error(
			"failed to connect to database",
			slog.String("error", err.Error()),
		)
		os.Exit(1)
	}

	// инициализируем репозиторий
	// передаем туда клиент базы данных
	repo := repository.New(lg, client)

	// инициализируем сервисы
	authService := auth.New(cfg, lg, repo)
	userService := user.New(lg, repo)
}
