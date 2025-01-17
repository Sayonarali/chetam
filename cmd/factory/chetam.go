package factory

import (
	"chetam/cfg"
	dbClient "chetam/internal/db/client"
	chetam "chetam/internal/service"
	"chetam/internal/service/auth"
	"github.com/google/wire"
	"log/slog"
)

var chSet = wire.NewSet(
	provideChetam,
	provideAuthService,
	provideChetamFetcher,
)

func provideChetam(
	lg *slog.Logger,
	as auth.Service,
	chf *dbClient.ChetamFetcher,
) chetam.Chetam {
	return chetam.New(lg, &as, chf)
}

func provideAuthService() (auth.Service, error) {
	c := auth.Config{}
	if err := cfg.Parse(&c); err != nil {
		return auth.Service{}, err
	}
	return auth.NewAuthService(c), nil
}

func provideChetamFetcher() (*dbClient.ChetamFetcher, error) {
	c := dbClient.Config{}
	if err := cfg.Parse(&c); err != nil {
		return nil, err
	}
	return dbClient.NewDBFetcher(c)
}
