package user

import (
	"chetam/internal/transport/repository"
	chetamApiv1 "chetam/pkg/chetamApi/v1"
	"log/slog"
)

type Service struct {
	lg   *slog.Logger
	repo *repository.Repository
}

func New(lg *slog.Logger, repo *repository.Repository) *Service {
	return &Service{
		lg:   lg,
		repo: repo,
	}
}

func (s *Service) GetUserByLogin(login string) (chetamApiv1.User, error) {
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
