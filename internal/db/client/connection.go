package client

import (
	"chetam/cfg"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Config struct {
	cfg.DB `envPrefix:"DB"`
}

type ChetamFetcher struct {
	c *sql.DB
}

func NewDBFetcher(cfg Config) (*ChetamFetcher, error) {
	connStr := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable",
		cfg.DB.Host, cfg.DB.Port, cfg.DB.User, cfg.DB.Name)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("db connection error %v", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("db connection error %v", err)
	}

	return &ChetamFetcher{
		c: db,
	}, nil
}
