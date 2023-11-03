package repository

import (
	"database/sql"

	"github.com/flpcastro/golang-kafka-impl/internal/entity"
)

type ProductRepositoryMySql struct {
	DB *sql.DB
}

func NewProductRepositoryMySql(db *sql.DB) *ProductRepositoryMySql {
	return &ProductRepositoryMySql{DB: db}
}

func (r *ProductRepositoryMySql) Create(product *entity.Product) error {
	_, err := r.DB.Exec("INSERT INTO products (id, name, price) values(?, ?, ?)", product.ID, product.Name, product.Price)

	if err != nil {
		return err
	}

	return nil
}

func (r *ProductRepositoryMySql) FindAll() ([]*entity.Product, error) {
	rows, err := r.DB.Query("SELECT id, name, price FROM products")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var products []*entity.Product
	for rows.Next() {
		var product entity.Product
		err = rows.Scan(&product.ID, &product.Name, &product.Price)
		if err != nil {
			return nil, err
		}

		products = append(products, &product)
	}

	return products, nil
}
