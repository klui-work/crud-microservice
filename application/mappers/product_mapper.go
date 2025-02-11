package mappers

import (
	"crud/application/dtos"
	"crud/domain/entities"
	"crud/helpers"

	"strconv"
)

func ProductMapper(product dtos.ProductDto) entities.Product {
	productPrice, err := strconv.ParseFloat(product.ProductPrice, 64)
	if err != nil {
		productPrice = 0
	}
	productStock, err := strconv.ParseUint(product.ProductStock, 10, 64)
	if err != nil {
		productStock = 0
	}
	return entities.Product{
		ProductId:    helpers.GenerateUUID(),
		ProductName:  product.ProductName,
		ProductPrice: productPrice,
		ProductStock: uint(productStock),
	}
}

func QueryProductMapper(query dtos.QueryProduct) entities.QueryProduct {
	return entities.QueryProduct{
		ProductId:   query.ProductId,
		ProductName: query.ProductName,
	}
}

func UpdateProductMapper(product dtos.ProducUpdatetDto) entities.Product {
	productPrice, err := strconv.ParseFloat(product.ProductPrice, 64)
	if err != nil {
		productPrice = 0
	}
	productStock, err := strconv.ParseUint(product.ProductStock, 10, 64)
	if err != nil {
		productStock = 0
	}
	return entities.Product{
		ProductId:    product.ProductId,
		ProductName:  product.ProductName,
		ProductPrice: productPrice,
		ProductStock: uint(productStock),
	}
}
