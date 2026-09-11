package orders

import (
	"context"
	"errors"
)

var (
	ErrInvalidQuantity   = errors.New("quantity must be greater than zero")
	ErrInsufficientStock = errors.New("insufficient stock")
)

type Service struct {
	orders   Repository
	products ProductClient
}

func NewService(orders Repository, productClient ProductClient) *Service {
	return &Service{
		orders:   orders,
		products: productClient,
	}
}

func (s *Service) GetAll() []Order {
	return s.orders.GetAll()
}

func (s *Service) Create(ctx context.Context, req CreateOrderRequest) (Order, error) {
	if req.Quantity <= 0 {
		return Order{}, ErrInvalidQuantity
	}

	product, err := s.products.GetProduct(ctx, req.ProductID)
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
