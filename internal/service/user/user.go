package user

import (
	"chetam/internal/service/repository"
	chetamApiv1 "chetam/pkg/chetamApi/v1"
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

func (s Service) GetUserByLogin(login string) (chetamApiv1.User, error) {
	user, err := s.repositoryKeeper.FindUserByLogin(login)
	if err != nil {
		s.lg.Warn(err.Error())
		return chetamApiv1.User{}, err
	}

	return chetamApiv1.User{
		Email:  &user.Email,
		Login:  &user.Login,
		UserId: &user.Id,
	}, nil
}
