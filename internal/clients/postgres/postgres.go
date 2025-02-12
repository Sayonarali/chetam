package postgres

import (
	"chetam/internal/cfg"
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
)

func NewClient(cfg *cfg.Config) (*pgx.Conn, error) {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.DB.User,
		cfg.DB.Password,
		cfg.DB.Host,
		cfg.DB.Port,
		cfg.DB.Name)

	db, err := pgx.Connect(context.Background(), dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping(context.Background())
	if err != nil {
		return nil, err
	}

	return db, nil
}
