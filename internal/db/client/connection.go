package client

import (
	"chetam/cfg"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Config struct {
	cfg.DB `envPrefix:"DB_"`
}

type ChetamFetcher struct {
	c *sql.DB
}

func NewDBFetcher(cfg Config) (*ChetamFetcher, error) {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		"postgres", 5432, "root", "root", "chetam")

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
