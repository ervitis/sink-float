//go:build wireinject
// +build wireinject

//go:generate go run github.com/google/wire/cmd/wire

package registry

import (
	"github.com/ervitis/sink-float/master/config"
	"github.com/google/wire"

	"github.com/ervitis/sink-float/master/handlers/grpc"
	"github.com/ervitis/sink-float/master/repository"
	"github.com/ervitis/sink-float/master/usecases"
)

var repositoriesSet = wire.NewSet(
	repository.NewMemcache,
)

var usecasesSet = wire.NewSet(
	repositoriesSet,
	usecases.NewSinkUseCase,
)

type GRPCHandlers struct {
	grpc.SinkFleetHandler
}

func NewGRPCHandlers(sinkFleetHandler grpc.SinkFleetHandler) GRPCHandlers {
	return GRPCHandlers{
		sinkFleetHandler,
	}
}

func HandleGRPCHandlerProvider(cfg config.AppConfig) GRPCHandlers {
	wire.Build(
		usecasesSet,
		NewGRPCHandlers,
	)
	return GRPCHandlers{}
}
