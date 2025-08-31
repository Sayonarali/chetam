package point

import (
	"chetam/internal/repository"
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
