package service

import (
	"crud/application/dtos"
	"crud/domain/entities"
)

type IProductService interface {
	CreateProductService(product dtos.ProductDto) error
	UpdateProductService(product dtos.ProductDto) error
	DeleteProductService(productId string) error
	GetProductByQueryService(query dtos.QueryProduct) ([]entities.Product, error)
}
