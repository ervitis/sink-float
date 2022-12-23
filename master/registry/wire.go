//go:build wireinject
// +build wireinject

//go:generate go run github.com/google/wire/cmd/wire

package registry

import "github.com/google/wire"

var repositoriesSet = wire.NewSet()

var usecasesSet = wire.NewSet()

type GRPCHandlers struct {
}

func NewGRPCHandlers() GRPCHandlers {
	return GRPCHandlers{}
}

func HandleGRPCHandlerProvider() GRPCHandlers {
	wire.Build(
		NewGRPCHandlers,
	)
	return GRPCHandlers{}
}
