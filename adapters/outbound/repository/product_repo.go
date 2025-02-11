package repository

import (
	"context"
	"fmt"

	"crud/domain/entities"
	"crud/domain/repository"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
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
	fmt.Println("product", product)
	_, err := p.collection.UpdateOne(context.Background(), bson.M{"product_id": product.ProductId}, bson.M{"$set": product})
	if err != nil {
		return err
	}
	return nil
}

func (p *ProductRepository) DeleteProduct(productId string) error {
	_, err := p.collection.DeleteOne(
		context.Background(),
		bson.M{"product_id": productId},
	)
	if err != nil {
		return err
	}
	return nil
}

func (p *ProductRepository) GetProductByQuery(query entities.QueryProduct) ([]entities.Product, error) {
	fmt.Println("query", query)
	filter := bson.M{}
	if query.ProductId != "" {
		filter["product_id"] = query.ProductId
	}
	if query.ProductName != "" {
		filter["product_name"] = query.ProductName
	}
	cursor, err := p.collection.Find(context.Background(), filter)
	if err != nil {
		return []entities.Product{}, err
	}
	var products []entities.Product
	if err := cursor.All(context.Background(), &products); err != nil {
		return []entities.Product{}, err
	}
	return products, nil
}

func NewProductRepository(db *mongo.Database) repository.IProductDomainRepo {
	c := &ProductRepository{
		collection: db.Collection("col_products"),
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := c.collection.Indexes().
		CreateOne(ctx, mongo.IndexModel{
			Keys: bson.D{{"product_id", 1}},
		})
	if err != nil {
		log.Println(err.Error())
	}
	return c
}
