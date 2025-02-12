//go:build wireinject
// +build wireinject

package factory

import (
	"chetam/internal/services/router"
	"github.com/google/wire"
)

func InitializeServer() (router.Router, error) {
	panic(
		wire.Build(
			srvSet,
			provideSlog,
		),
	)
}
