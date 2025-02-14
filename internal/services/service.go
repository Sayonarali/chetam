package services

import (
	"chetam/internal/auth"
	"log/slog"
)

type Services struct {
	logger *slog.Logger
	auth   *auth.Auth
}

func New(logger *slog.Logger, auth *auth.Auth) *Services {
	return &Services{
		logger: logger,
		auth:   auth,
	}
}
