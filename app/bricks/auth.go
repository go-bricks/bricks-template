package bricks

import (
	"github.com/go-bricks/bricks/constructors"
	"go.uber.org/fx"
)

func AuthFxOptions() fx.Option {
	return fx.Options(
		fx.Provide(constructors.DefaultJWTTokenExtractor),
	)
}
