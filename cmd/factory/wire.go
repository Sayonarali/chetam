//go:build wireinject
// +build wireinject

package factory

import (
	"chetam/internal/service/router"
	"github.com/google/wire"
)

func InitializeRouter() (router.Router, error) {
	panic(
		wire.Build(
			chSet,
			provideSlog,
		),
	)
}
