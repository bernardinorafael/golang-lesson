package entity

import (
	"github.com/google/uuid"
)

type ProductRepository interface {
	Create(product *Product) error
	FindAll() ([]*Product, error)
}

type Product struct {
	ID    string
	Name  string
	Price float64
}

func CreateNewProduct(name string, price float64) *Product {
	return &Product{
		ID:    uuid.NewString(),
		Name:  name,
		Price: price,
	}
}
