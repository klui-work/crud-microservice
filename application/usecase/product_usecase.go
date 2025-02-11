package usecase

import (
	"crud/application/dtos"
	"crud/domain/entities"
	"crud/domain/repository"
	"crud/domain/service"
)

type productService struct {
	productRepo repository.IProductDomainRepo
}

func (p *productService) CreateProductService(product dtos.ProductDto) error {
	return nil
}

func (p *productService) UpdateProductService(product dtos.ProductDto) error {
	return nil
}

func (p *productService) DeleteProductService(productId string) error {
	return nil
}

func (p *productService) GetProductByQueryService(query dtos.QueryProduct) ([]entities.Product, error) {
	return nil, nil
}

func NewProductService(
	productRepo repository.IProductDomainRepo,
) service.IProductService {
	return &productService{
		productRepo: productRepo,
	}
}
