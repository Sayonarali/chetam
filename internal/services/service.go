package services

import (
	"chetam/internal/auth"
	"chetam/internal/services/point"
	"chetam/internal/services/route"
	"log/slog"
)

type Services struct {
	logger *slog.Logger
	Auth   *auth.Auth
	Route  *route.Service
	Point  *point.Service
}

func New(
	logger *slog.Logger,
	auth *auth.Auth,
	route *route.Service,
	point *point.Service) *Services {
	return &Services{
		logger: logger,
		Auth:   auth,
		Route:  route,
		Point:  point,
	}
}
