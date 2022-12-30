package main

import (
	"github.com/alecthomas/kong"
	"github.com/go-bricks/bricks-template/app/bricks"
	"github.com/go-bricks/bricks/providers"
	"go.uber.org/fx"
)

var CLI struct {
	Config struct {
		Path            string   `arg:"" required:"" help:"Path to config file." type:"existingfile"`
		AdditionalFiles []string `optional:"" help:"Additional configuration files to merge, comma separated" type:"existingfile"`
	} `cmd:"" help:"Path to config file."`
}

func main() {
	ctx := kong.Parse(&CLI, kong.UsageOnError())
	switch cmd := ctx.Command(); cmd {
	case "config <path>":
		app := createApplication(CLI.Config.Path, CLI.Config.AdditionalFiles)
		app.Run()
	default:
		ctx.Fatalf("unknown option %s", cmd)
	}
}

func createApplication(configFilePath string, additionalFiles []string) *fx.App {
	return fx.New(
		// fx.NopLogger, // remove fx debug
		bricks.ViperFxOption(configFilePath, additionalFiles...), // Configuration map
		bricks.LoggerFxOption(),                                  // Logger
		bricks.TracerFxOption(),                                  // Jaeger tracing
		bricks.PrometheusFxOption(),                              // Prometheus
		bricks.HttpClientFxOptions(),
		bricks.HttpServerFxOptions(),
		bricks.AuthFxOptions(),
		bricks.InternalHttpHandlersFxOptions(),
		// Tutorial service dependencies
		bricks.ServiceAPIsAndOtherDependenciesFxOption(), // register tutorial APIs
		// This one invokes all the above
		providers.BuildBricksWebServiceFxOption(), // http server invoker
	)
}
