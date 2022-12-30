package bricks

import (
	"os"

	"github.com/go-bricks/bjaeger"
	"github.com/go-bricks/bzerolog"
	"github.com/go-bricks/bricks/interfaces/cfg"
	"github.com/go-bricks/bricks/interfaces/log"
	"github.com/go-bricks/bricks/providers"
	"go.uber.org/fx"
)

func LoggerFxOption() fx.Option {
	return fx.Options(
		fx.Provide(zeroLogBuilder),
		providers.LoggerFxOption(),
		providers.LoggerGRPCIncomingContextExtractorFxOption(),
		bjaeger.TraceInfoContextExtractorFxOption(),
	)
}

func zeroLogBuilder(config cfg.Config) log.Builder {
	builder := bzerolog.Builder().IncludeCaller()
	if config.Get("server.logger.console").Bool() {
		builder = builder.
			SetWriter(bzerolog.ConsoleWriter(os.Stderr))
	}
	return builder
}
