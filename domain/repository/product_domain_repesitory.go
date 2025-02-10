package repository

import (
	"crud/domain/entities"
)

type IProductDomainRepo interface {
	CreateProduct(product entities.Product) error
	UpdateProduct(product entities.Product) error
	DeleteProduct(productId string) error
	GetProductByQuery(query entities.QueryProduct) ([]entities.Product, error)
}
