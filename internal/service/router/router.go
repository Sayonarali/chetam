package router

import (
	"chetam/internal/service/auth"
	"chetam/internal/service/user"
	"log/slog"
)

type Router struct {
	lg          *slog.Logger
	userService *user.Service
	authService *auth.Service
}

func New(
	lg *slog.Logger,
	us *user.Service,
	as *auth.Service,
) Router {
	if as == nil || lg == nil {
		panic("nil argument in new")
	}
	return Router{
		lg:          lg,
		userService: us,
		authService: as,
	}
}
