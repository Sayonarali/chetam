package repository

import (
	"log/slog"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
)

type Repository struct {
	logger *slog.Logger
	db     *pgx.Conn
}

func New(logger *slog.Logger, db *pgx.Conn) *Repository {
	return &Repository{
		logger: logger,
		db:     db,
	}
}

var qb = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)
