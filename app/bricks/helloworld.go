package bricks

import (
	"context"

	helloworld "github.com/go-bricks/bricks-template/api"
	"github.com/go-bricks/bricks-template/app/controllers"
	"github.com/go-bricks/bricks-template/app/services"
	"github.com/go-bricks/bricks-template/app/validations"
	serverInt "github.com/go-bricks/bricks/interfaces/http/server"
	"github.com/go-bricks/bricks/providers/groups"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go.uber.org/fx"
	"google.golang.org/grpc"
)

type workshopServiceDeps struct {
	fx.In

	// API Implementations, "Register" them as GRPCServiceAPI
	Helloworld helloworld.GreeterServer
}

func ServiceAPIsAndOtherDependenciesFxOption() fx.Option {
	return fx.Options(
		// GRPC Service APIs registration
		fx.Provide(fx.Annotated{
			Group:  groups.GRPCServerAPIs,
			Target: serviceGRPCServiceAPIs,
		}),
		// GRPC Gateway Generated Handlers registration
		fx.Provide(fx.Annotated{
			Group:  groups.GRPCGatewayGeneratedHandlers + ",flatten", // "flatten" does this [][]serverInt.GRPCGatewayGeneratedHandlers -> []serverInt.GRPCGatewayGeneratedHandlers
			Target: serviceGRPCGatewayHandlers,
		}),
		// All other tutorial dependencies
		serviceDependencies(),
	)
}

func serviceGRPCServiceAPIs(deps workshopServiceDeps) serverInt.GRPCServerAPI {
	return func(srv *grpc.Server) {
		helloworld.RegisterGreeterServer(srv, deps.Helloworld)
		// Any additional gRPC Implementations should be called here
	}
}

func serviceGRPCGatewayHandlers() []serverInt.GRPCGatewayGeneratedHandlers {
	return []serverInt.GRPCGatewayGeneratedHandlers{
		// Register service REST API
		func(mux *runtime.ServeMux, localhostEndpoint string) error {
			return helloworld.RegisterGreeterHandlerFromEndpoint(context.Background(), mux, localhostEndpoint, []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())})
		},
		// Any additional gRPC gateway registrations should be called here
	}
}

func serviceDependencies() fx.Option {
	return fx.Provide(
		services.CreateHelloworldService,
		controllers.CreateHelloworldController,
		validations.CreateHelloworldValidations,
	)
}
