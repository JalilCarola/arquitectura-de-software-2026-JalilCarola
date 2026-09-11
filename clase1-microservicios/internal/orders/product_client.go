package orders

import (
	"context"
	"errors"
)

var ErrProductNotFound = errors.New("product not found")

type Product struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
	Stock int     `json:"stock"`
}

type ProductClient interface {
	GetProduct(ctx context.Context, id string) (Product, error)
}
