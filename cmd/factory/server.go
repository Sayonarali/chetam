package factory

import (
	"chetam/internal/http/srv"
	"chetam/internal/services"
	"github.com/google/wire"
	"log/slog"
)

var srvSet = wire.NewSet(
	provideServer,
	provideAuthService,
	provideUserService,
	provideChetamFetcher,
	provideRepositoryKeeper,
)

func provideServer(
	lg *slog.Logger,
	service *services.Service,
) srv.Server {
	return srv.New(lg, service)
}

//
//func provideUserService(
//	rk repository.Keeper,
//	lg *slog.Logger,
//) (user.Service, error) {
//	return user.NewUserService(rk, lg), nil
//}
//
//func provideAuthService(
//	rk repository.Keeper,
//	lg *slog.Logger,
//) (auth.Service, error) {
//	c := auth.Config{}
//	if err := cfg.Parse(&c); err != nil {
//		return auth.Service{}, err
//	}
//	return auth.NewAuthService(c, rk, lg), nil
//}
//
//func provideRepositoryKeeper(chf *dbClient.ChetamFetcher) (repository.Keeper, error) {
//	return repository.NewRepositoryKeeper(chf), nil
//}
//
//func provideChetamFetcher() (*dbClient.ChetamFetcher, error) {
//	c := dbClient.Config{}
//	if err := cfg.Parse(&c); err != nil {
//		return nil, err
//	}
//	return dbClient.NewDBFetcher(c)
//}
