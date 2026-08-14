package orders

import "strconv"

type Repository interface {
	GetAll() []Order
	Save(order Order) Order
}

type InMemoryRepository struct {
	orders []Order
	nextID int
}

func NewRepository() *InMemoryRepository {
	return &InMemoryRepository{
		orders: make([]Order, 0),
		nextID: 1,
	}
}

func (r *InMemoryRepository) GetAll() []Order {
	result := make([]Order, len(r.orders))
	copy(result, r.orders)
	return result
}

func (r *InMemoryRepository) Save(order Order) Order {
	order.ID = strconv.Itoa(r.nextID)
	r.nextID++
	r.orders = append(r.orders, order)
	return order
}
