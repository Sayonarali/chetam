package user

import (
	chetamApiv1 "chetam/internal/server/client/v1"
	"chetam/internal/transport/repository"
	"log/slog"
)

type UserService struct {
	lg   *slog.Logger
	repo *repository.Repository
}

func New(lg *slog.Logger, repo *repository.Repository) *UserService {
	return &UserService{
		lg:   lg,
		repo: repo,
	}
}

func (s *UserService) GetUserByLogin(login string) (chetamApiv1.User, error) {
	user, err := s.repo.FindUserByLogin(login)
	if err != nil {
		s.lg.Warn("user not found",
			slog.String("error", err.Error()),
		)

		return chetamApiv1.User{}, err
	}

	return chetamApiv1.User{
		Email:  &user.Email,
		Login:  &user.Login,
		UserId: &user.Id,
	}, nil
}
