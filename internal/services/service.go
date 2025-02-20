package services

import (
	"chetam/internal/services/auth"
	"log/slog"
)

type Services struct {
	logger *slog.Logger
	Auth   *auth.Auth
}

func New(logger *slog.Logger, auth *auth.Auth) *Services {
	return &Services{
		logger: logger,
		Auth:   auth,
	}
}
