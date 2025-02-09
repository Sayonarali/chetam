package service

import (
	"chetam/internal/auth"
	"chetam/internal/service/user"
)

type Service struct {
	userService *user.Service
	authService *auth.Service
}

func New(userService *user.Service, authService *auth.Service) *Service {
	if userService == nil || authService == nil {
		panic("nil argument in new")
	}
	return &Service{
		userService: userService,
		authService: authService,
	}
}
