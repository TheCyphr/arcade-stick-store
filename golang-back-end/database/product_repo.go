package database

import (
	"database/sql"
	"fmt"
	"github.com/arcade-stick-store/models"
)

type ProductRepo struct {
	sql *sql.DB
}

func NewRepo(sql *sql.DB) *ProductRepo {
	return &ProductRepo{sql: sql}
}

func (pr *ProductRepo) Products() ([]models.Product, error) {
	var products []models.Product

	rows, err := pr.sql.Query("SELECT * FROM products")
	if err != nil {
		return nil, fmt.Errorf("error querying products table: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var prod models.Product
		if err := rows.Scan(&prod.Id, &prod.Name, &prod.Description, &prod.Brand, &prod.Type, &prod.Price, &prod.StockQuantity, &prod.Image); err != nil {
			return nil, fmt.Errorf("error retrieving product: %v", err)
		}
		products = append(products, prod)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error retrieving product: %v", err)
	}
	return products, nil
}
