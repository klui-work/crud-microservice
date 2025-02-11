package validation

import (
	"crud/application/dtos"
	"errors"
)

func ValidationCreateProduct(product dtos.ProductDto) error {
	if product.ProductName == "" {
		return errors.New("product name is required")
	}
	if product.ProductPrice == "" {
		return errors.New("product price is required")
	}
	if product.ProductStock == "" {
		return errors.New("product stock is required")
	}
	return nil
}

func ValidationUpdateProduct(product dtos.ProducUpdatetDto) error {
	if product.ProductId == "" {
		return errors.New("product id is required")
	}
	return nil
}

func ValidationDeleteProduct(productId string) error {
	if productId == "" {
		return errors.New("product id is required")
	}
	return nil
}
