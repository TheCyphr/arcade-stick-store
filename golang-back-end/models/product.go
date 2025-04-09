package models

type Product struct {
	Id            int
	Name          string
	Description   string
	Brand         string
	Type          string
	Price         float64
	StockQuantity int
	Image         string
}
