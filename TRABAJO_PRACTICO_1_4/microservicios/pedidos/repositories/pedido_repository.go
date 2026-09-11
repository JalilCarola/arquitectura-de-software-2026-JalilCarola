package repositories

import (
	"strconv"
	"sync"
)

type Pedido struct {
	ID         string `json:"id"`
	ClienteID  string `json:"cliente_id"`
	ProductoID string `json:"producto_id"`
}

type PedidoRepository struct {
	mu      sync.Mutex
	pedidos []Pedido
	nextID  int
}

func NewPedidoRepository() *PedidoRepository {
	return &PedidoRepository{
		pedidos: make([]Pedido, 0),
		nextID:  1,
	}
}

func (r *PedidoRepository) Guardar(clienteID, productoID string) Pedido {
	r.mu.Lock()
	defer r.mu.Unlock()

	pedido := Pedido{
		ID:         "PED-" + strconv.Itoa(r.nextID),
		ClienteID:  clienteID,
		ProductoID: productoID,
	}
	r.pedidos = append(r.pedidos, pedido)
	r.nextID++

	return pedido
}
