package services

import (
	"errors"
	"log"

	"pedidos/messaging"
	"pedidos/repositories"
)

var ErrProductoInexistente = errors.New("producto inexistente")

type PedidoService struct {
	pedidos   *repositories.PedidoRepository
	productos *ProductoService
	publisher *messaging.RabbitMQPublisher
}

func NewPedidoService(pedidos *repositories.PedidoRepository, productos *ProductoService, publisher *messaging.RabbitMQPublisher) *PedidoService {
	return &PedidoService{
		pedidos:   pedidos,
		productos: productos,
		publisher: publisher,
	}
}

func (s *PedidoService) Confirmar(clienteID, productoID string) (repositories.Pedido, error) {
	if !s.productos.Existe(productoID) {
		return repositories.Pedido{}, ErrProductoInexistente
	}

	pedido := s.pedidos.Guardar(clienteID, productoID)

	if err := s.publisher.PublicarPedidoConfirmado(pedido.ID, pedido.ClienteID, pedido.ProductoID); err != nil {
		// El pedido ya quedó confirmado y guardado; solo avisamos que el evento no salió.
		log.Printf("[EVENTO] No se pudo publicar pedido.confirmado para %s: %v", pedido.ID, err)
	}

	return pedido, nil
}
