package entities

type Product struct {
	ProductId    string  `bson:"product_id"`
	ProductName  string  `bson:"product_name"`
	ProductPrice float64 `bson:"product_price"`
	ProductStock uint    `bson:"product_stock"`
}

type QueryProduct struct {
	ProductId   string `bson:"product_id"`
	ProductName string `bson:"product_name"`
}
