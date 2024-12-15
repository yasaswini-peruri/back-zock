package services

import (
	"database/sql"
	"fmt"
	"productmanagement/models"

	_ "github.com/lib/pq"
)

var db *sql.DB

func InitDB() {
	connStr := "user=postgres password=password dbname=product_management sslmode=disable"
	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
}

func CreateProduct(product *models.Product) error {
	query := `INSERT INTO products (user_id, product_name, product_description, product_images, product_price)
			  VALUES ($1, $2, $3, $4, $5)`
	_, err := db.Exec(query, product.UserID, product.ProductName, product.ProductDescription, product.ProductImages, product.ProductPrice)
	return err
}

func GetProductByID(id string) (*models.Product, error) {
	query := `SELECT * FROM products WHERE id = $1`
	row := db.QueryRow(query, id)
	product := &models.Product{}
	err := row.Scan(&product.ID, &product.UserID, &product.ProductName, &product.ProductDescription, &product.ProductImages, &product.ProductPrice, &product.CreatedAt)
	return product, err
}

func GetProductsByUserID(userID string) ([]models.Product, error) {
	query := `SELECT * FROM products WHERE user_id = $1`
	rows, err := db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []models.Product
	for rows.Next() {
		product := models.Product{}
		if err := rows.Scan(&product.ID, &product.UserID, &product.ProductName, &product.ProductDescription, &product.ProductImages, &product.ProductPrice, &product.CreatedAt); err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}
