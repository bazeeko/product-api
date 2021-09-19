package data

import (
	"database/sql"
	"fmt"
)

// Product defines the structure for an API product
type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

type ProductsRepository struct {
	Database *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductsRepository {
	return &ProductsRepository{Database: db}
}

// GetProducts returns all products from the database
func (repo *ProductsRepository) GetProducts() (products []Product, err error) {
	rows, err := repo.Database.Query(`SELECT id, name, description, price FROM products`)
	if rows == nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		product := Product{}
		err = rows.Scan(&product.ID, &product.Name, &product.Description, &product.Price)
		if err != nil {
			return
		}
		products = append(products, product)
	}
	return
}

// GetProductById returns a single product which matches the id from the
// database.
func (repo *ProductsRepository) GetProductById(id int) (product Product, err error) {
	err = repo.Database.QueryRow(`SELECT id, name, description, price FROM products WHERE id = $1`, id).
		Scan(&product.ID, &product.Name, &product.Description, &product.Price)
	return
}

// AddProduct adds a new product to the database
func (repo *ProductsRepository) AddProduct(product *Product) (err error) {
	_, err = repo.Database.Exec(`INSERT INTO products (name, description, price) VALUES ($1, $2, $3)`,
		product.Name, product.Description, product.Price)
	if err != nil {
		fmt.Println(err)
	}
	return
}

// UpdateProduct replaces a product in the database with the given
// item.
func (repo *ProductsRepository) UpdateProduct(values Product) (err error) {
	_, err = repo.Database.Exec(`UPDATE products SET name = $1, description = $2, price = $3 WHERE id = $4`, values.Name, values.Description, values.Price, values.ID)
	return
}

// DeleteProduct deletes a product from the database
func (repo *ProductsRepository) DeleteProduct(id int) (err error) {
	_, err = repo.Database.Exec(`DELETE FROM products WHERE id = $1`, id)
	return
}
