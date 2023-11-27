package repository

import (
	"database/sql"

	"github.com/bernardinorafael/go-mensageria/internal/entity"
)

type ProductRepositoryDB struct {
	DB *sql.DB
}

func (r *ProductRepositoryDB) FindAll() ([]*entity.Product, error) {
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

func (r *ProductRepositoryDB) Create(p *entity.Product) error {
	_, err := r.DB.Exec(
		"INSERT INTO products (id, name, price) values (?,?,?)",
		p.ID, p.Name, p.Price,
	)

	if err != nil {
		return err
	}

	return nil
}

func CreateNewProductRepository(db *sql.DB) *ProductRepositoryDB {
	return &ProductRepositoryDB{db}
}
