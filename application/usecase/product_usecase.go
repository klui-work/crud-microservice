package usecase

import (
	"crud/application/dtos"
	"crud/application/mappers"
	"crud/domain/entities"
	"crud/domain/repository"
	"crud/domain/service"
)

type productService struct {
	productRepo repository.IProductDomainRepo
}

func (p *productService) CreateProductService(product dtos.ProductDto) error {
	productEntity := mappers.ProductMapper(product)
	return p.productRepo.CreateProduct(productEntity)
}

func (p *productService) UpdateProductService(product dtos.ProducUpdatetDto) error {
	productEntity := mappers.UpdateProductMapper(product)
	return p.productRepo.UpdateProduct(productEntity)
}

func (p *productService) DeleteProductService(productId string) error {
	return p.productRepo.DeleteProduct(productId)
}

func (p *productService) GetProductByQueryService(query dtos.QueryProduct) ([]entities.Product, error) {
	queryEntity := mappers.QueryProductMapper(query)
	results, err := p.productRepo.GetProductByQuery(queryEntity)
	if err != nil {
		return []entities.Product{}, err
	}
	return results, nil
}

func NewProductService(
	productRepo repository.IProductDomainRepo,
) service.IProductService {
	return &productService{
		productRepo: productRepo,
	}
}
