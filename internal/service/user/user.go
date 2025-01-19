package user

import (
	"chetam/internal/service/repository"
	"log/slog"
)

type Service struct {
	lg               *slog.Logger
	repositoryKeeper repository.Keeper
}

func NewUserService(rk repository.Keeper, lg *slog.Logger) Service {
	return Service{
		lg:               lg,
		repositoryKeeper: rk,
	}
}

func (s Service) GetUser() {

}
