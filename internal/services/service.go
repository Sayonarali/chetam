package services

import (
	"log/slog"
)

type Services struct {
	logger *slog.Logger
}

func New(logger *slog.Logger) *Services {
	return &Services{
		logger: logger,
	}
}
