package main

import (
	"chetam/internal/config"
	"chetam/internal/db/client/postgres"
	"chetam/internal/server"
	"chetam/internal/server/handlers"
	"chetam/internal/services"
	"chetam/internal/services/auth"
	"chetam/internal/transport/repository"
	"chetam/pkg/logger"
	"fmt"
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
	fmt.Println(repo)
	a := auth.New(cfg, lg, repo)
	s := services.New(lg, a)

	handler := handlers.New(lg, s)
	srv := server.New(lg, cfg, handler)

	srv.Run()
}
