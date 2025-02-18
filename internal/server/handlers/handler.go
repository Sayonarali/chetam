package handlers

import (
	"chetam/internal/services"
	"log/slog"
)

type ServerHandler struct {
	services *services.Services
	lg       *slog.Logger
}

func New(logger *slog.Logger, services *services.Services) *ServerHandler {
	return &ServerHandler{
		lg:       logger,
		services: services,
	}
}
