//go:build wireinject
// +build wireinject

package injector

import (
	"crud/adapters/inbound"
	"crud/adapters/inbound/api/controller"

	"github.com/google/wire"
)

func ProvideAdapter(
	ProductCtl controller.IProductController,
) *inbound.InitialAdapter {
	return inbound.NewInitialAdapter(ProductCtl)
}

var ProviderSet = wire.NewSet(
	ProvideAdapter,
	RepositoryProviderSet,
	ControllerProvideSet,
	ServiceProviderSet,
)

func InitializeAdapter() *inbound.InitialAdapter {
	wire.Build(ProviderSet)
	return &inbound.InitialAdapter{}
}
