package dtos

type ProductDto struct {
	ProductName  string `json:"product_name"`
	ProductStock string `json:"product_stock"`
	ProductPrice string `json:"product_price"`
}

type ProducUpdatetDto struct {
	ProductId    string `json:"product_id"`
	ProductName  string `json:"product_name"`
	ProductStock string `json:"product_stock"`
	ProductPrice string `json:"product_price"`
}

type QueryProduct struct {
	ProductId   string `json:"product_id"`
	ProductName string `json:"product_name"`
}
