package orders

import (
	"errors"

	"clase1/starter/internal/products"
)

var (
	ErrInvalidQuantity   = errors.New("quantity must be greater than zero")
	ErrInsufficientStock = errors.New("insufficient stock")
)

type Service struct {
	orders   Repository
	products products.Repository
}

func NewService(orders Repository, productRepository products.Repository) *Service {
	return &Service{
		orders:   orders,
		products: productRepository,
	}
}

func (s *Service) GetAll() []Order {
	return s.orders.GetAll()
}

func (s *Service) Create(req CreateOrderRequest) (Order, error) {
	if req.Quantity <= 0 {
		return Order{}, ErrInvalidQuantity
	}

	product, err := s.products.GetByID(req.ProductID)
	if err != nil {
		return Order{}, err
	}

	if product.Stock < req.Quantity {
		return Order{}, ErrInsufficientStock
	}

	order := Order{
		ProductID: req.ProductID,
		Quantity:  req.Quantity,
		Total:     product.Price * float64(req.Quantity),
	}

	return s.orders.Save(order), nil
}
