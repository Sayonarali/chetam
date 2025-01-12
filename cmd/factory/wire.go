//go:build wireinject
// +build wireinject

package factory

import (
	chetam "chetam/internal/service"
	"github.com/google/wire"
)

func InitializeChetam() (chetam.Chetam, error) {
	panic(
		wire.Build(
			chSet,
			provideSlog,
		),
	)
}
