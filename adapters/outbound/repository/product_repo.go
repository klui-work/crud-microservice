package repository

import (
	"context"
	// "crud/application/dtos"
	"crud/domain/entities"
	"crud/domain/repository"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

type ProductRepository struct {
	collection *mongo.Collection
}

func (p *ProductRepository) CreateProduct(product entities.Product) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err := p.collection.InsertOne(ctx, product)
	if err != nil {
		return err
	}
	return nil
}

func (p *ProductRepository) UpdateProduct(product entities.Product) error {
	return nil
}

func (p *ProductRepository) DeleteProduct(productId string) error {
	return nil
}

func (p *ProductRepository) GetProductByQuery(query entities.QueryProduct) ([]entities.Product, error) {
	return []entities.Product{}, nil
}
func NewProductRepository(db *mongo.Database) repository.IProductDomainRepo {
	return &ProductRepository{
		collection: db.Collection("col_products"),
	}

}
