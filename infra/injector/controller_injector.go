package injector

import (
	"crud/adapters/inbound/api/controller"
	"crud/domain/service"

	"github.com/google/wire"
)

func ProvideProductServiceController(service service.IProductService) controller.IProductController {
	return controller.NewProductController(service)
}

var ControllerProvideSet = wire.NewSet(
	ProvideProductServiceController,
)
