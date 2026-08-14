package products

import "errors"

var ErrProductNotFound = errors.New("product not found")

type Repository interface {
	GetAll() []Product
	GetByID(id string) (Product, error)
}

type InMemoryRepository struct {
	products []Product
}

func NewRepository() *InMemoryRepository {
	return &InMemoryRepository{
		products: []Product{
			{ID: "1", Name: "Notebook", Price: 1200, Stock: 10},
			{ID: "2", Name: "Mouse", Price: 25, Stock: 30},
			{ID: "3", Name: "Keyboard", Price: 60, Stock: 20},
		},
	}
}

func (r *InMemoryRepository) GetAll() []Product {
	result := make([]Product, len(r.products))
	copy(result, r.products)
	return result
}

func (r *InMemoryRepository) GetByID(id string) (Product, error) {
	for _, product := range r.products {
		if product.ID == id {
			return product, nil
		}
	}

	return Product{}, ErrProductNotFound
}
