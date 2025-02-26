package main

import (
	"chetam/internal/auth"
	"chetam/internal/config"
	"chetam/internal/db/client/postgres"
	"chetam/internal/server"
	"chetam/internal/services"
	"chetam/internal/services/user"
	"chetam/internal/transport/repository"
	"chetam/pkg/logger"
	"log/slog"
	"os"
)

func main() {
	lg := logger.New()
	slog.SetDefault(lg)

	cfg, err := config.Load()
	if err != nil {
		lg.Error(
			"failed to load config",
			slog.String("error", err.Error()),
		)
	}

	client, err := postgres.NewClient(cfg)
	if err != nil {
		lg.Error("failed to connect to database",
			slog.String("error", err.Error()))

		os.Exit(1)
	}

	repo := repository.New(lg, client)
	a := auth.New(cfg, lg, repo)
	userService := user.New(lg, repo)
	s := services.New(lg, a, userService)

	srv := server.New(lg, cfg, s)

	srv.Run()
}
