package chetam

import (
	dbClient "chetam/internal/db/client"
	"chetam/internal/service/auth"
	"github.com/gorilla/mux"
	"log/slog"
)

type Chetam struct {
	lg            *slog.Logger
	authService   *auth.Service
	chetamFetcher *dbClient.ChetamFetcher
}

func New(
	lg *slog.Logger,
	as *auth.Service,
	chf *dbClient.ChetamFetcher,
) Chetam {
	if as == nil || chf == nil || lg == nil {
		panic("nil argument in new")
	}
	return Chetam{
		lg:            lg,
		authService:   as,
		chetamFetcher: chf,
	}
}

func (c Chetam) Execute() {
	r := mux.NewRouter()

	r.HandleFunc("/api/v1/auth/register", c.authService.Register).Methods("POST")
	r.HandleFunc("/api/v1/auth/login", c.authService.Login).Methods("POST")

}
