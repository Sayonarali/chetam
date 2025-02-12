package services

import (
	"chetam/internal/services/auth"
	"chetam/internal/services/user"
)

type Services struct {
	user *user.Service
	auth *auth.Service
}

func New(userService *user.Service, authService *auth.Service) *Services {
	return &Services{
		user: userService,
		auth: authService,
	}
}
