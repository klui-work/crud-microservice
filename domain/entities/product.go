package entities

type Product struct {
	ProductId    string
	ProductName  string
	ProductPrice float64
	ProductStock uint
}

type QueryProduct struct {
	ProductId   string
	ProductName string
}
