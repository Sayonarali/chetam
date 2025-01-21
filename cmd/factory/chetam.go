package factory

import (
	"chetam/cfg"
	dbClient "chetam/internal/db/client"
	"chetam/internal/service/auth"
	"chetam/internal/service/repository"
	"chetam/internal/service/router"
	"chetam/internal/service/user"
	"github.com/google/wire"
	"log/slog"
)

var chSet = wire.NewSet(
	provideRouterHandler,
	provideAuthService,
	provideChetamFetcher,
	provideRepositoryKeeper,
	provideUserService,
)

func provideRouterHandler(
	lg *slog.Logger,
	as auth.Service,
	us user.Service,
) router.Router {
	return router.New(lg, &us, &as)
}

func provideUserService(
	rk repository.Keeper,
	lg *slog.Logger,
) (user.Service, error) {
	return user.NewUserService(rk, lg), nil
}

func provideAuthService(
	rk repository.Keeper,
	lg *slog.Logger,
) (auth.Service, error) {
	c := auth.Config{}
	if err := cfg.Parse(&c); err != nil {
		return auth.Service{}, err
	}
	return auth.NewAuthService(c, rk, lg), nil
}

func provideChetamFetcher() (*dbClient.ChetamFetcher, error) {
	c := dbClient.Config{}
	if err := cfg.Parse(&c); err != nil {
		return nil, err
	}
	return dbClient.NewDBFetcher(c)
}

func provideRepositoryKeeper(chf *dbClient.ChetamFetcher) (repository.Keeper, error) {
	return repository.NewRepositoryKeeper(chf), nil
}
