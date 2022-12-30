package bricks

import (
	"github.com/go-bricks/bprometheus"
	"github.com/go-bricks/bricks/interfaces/cfg"
	confkeys "github.com/go-bricks/bricks/interfaces/cfg/keys"
	"github.com/go-bricks/bricks/interfaces/monitor"
	"github.com/go-bricks/bricks/providers"
	"go.uber.org/fx"
)

// PrometheusFxOption registers prometheus
func PrometheusFxOption() fx.Option {
	return fx.Options(
		providers.MonitorFxOption(),
		providers.MonitorGRPCInterceptorFxOption(),
		bprometheus.PrometheusInternalHandlerFxOption(),
		fx.Provide(PrometheusBuilder),
	)
}

// PrometheusBuilder returns a monitor.Builder that is implemented by Prometheus
func PrometheusBuilder(cfg cfg.Config) monitor.Builder {
	name := cfg.Get(confkeys.ApplicationName).String()
	return bprometheus.Builder().SetNamespace(name)
}
