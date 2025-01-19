//go:build wireinject
// +build wireinject

package factory

import (
	router "chetam/internal/service"
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
