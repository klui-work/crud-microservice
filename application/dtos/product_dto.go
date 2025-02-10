package dtos

type ProductDto struct {
	ProductName  string  `json:"product_name"`
	ProductStock uint    `json:"product_stock"`
	ProductPrice float64 `json:"product_price"`
}

type QueryProduct struct {
	ProductId   string `json:"product_id"`
	ProductName string `json:"product_name"`
}
