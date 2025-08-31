package main

import (
	"chetam/internal/auth"
	"chetam/internal/client"
	"chetam/internal/config"
	"chetam/internal/repository"
	"chetam/internal/server"
	"chetam/internal/services"
	"chetam/internal/services/point"
	"chetam/internal/services/route"
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

	cl, err := client.NewClient(cfg)
	if err != nil {
		lg.Error("failed to connect to database",
			slog.String("error", err.Error()))

		os.Exit(1)
	}

	repo := repository.New(lg, cl)
	a := auth.New(cfg, lg, repo)
	routes := route.New(lg, repo)
	points := point.New(lg, repo)
	s := services.New(lg, a, routes, points)

	srv := server.New(lg, cfg, s)

	srv.Run()
}
