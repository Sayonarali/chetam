package services

import (
	"chetam/internal/auth"
	"chetam/internal/services/user"
	"log/slog"
)

type Services struct {
	logger *slog.Logger
	Auth   *auth.Auth
	User   *user.UserService
}

func New(logger *slog.Logger, auth *auth.Auth, user *user.UserService) *Services {
	return &Services{
		logger: logger,
		Auth:   auth,
		User:   user,
	}
}
