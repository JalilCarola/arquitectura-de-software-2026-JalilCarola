package orders

type Order struct {
	ID        string  `json:"id"`
	ProductID string  `json:"productId"`
	Quantity  int     `json:"quantity"`
	Total     float64 `json:"total"`
}

type CreateOrderRequest struct {
	ProductID string `json:"productId"`
	Quantity  int    `json:"quantity"`
}
