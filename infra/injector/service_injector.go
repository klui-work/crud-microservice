package injector

import (
	"crud/application/usecase"
	"crud/domain/repository"
	"crud/domain/service"

	"github.com/google/wire"
)

func ProductService(
	productRepo repository.IProductDomainRepo,
) service.IProductService {
	return usecase.NewProductService(
		productRepo,
	)
}

var ServiceProviderSet = wire.NewSet(
	ProductService,
)
