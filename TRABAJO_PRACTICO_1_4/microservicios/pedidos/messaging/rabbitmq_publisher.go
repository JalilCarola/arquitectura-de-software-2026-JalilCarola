package messaging

import (
	"context"
	"encoding/json"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

const queueName = "pedidos-confirmados"

type PedidoConfirmado struct {
	Tipo       string `json:"tipo"`
	PedidoID   string `json:"pedido_id"`
	ClienteID  string `json:"cliente_id"`
	ProductoID string `json:"producto_id"`
}

type RabbitMQPublisher struct {
	channel *amqp.Channel
}

func NewRabbitMQPublisher(amqpURI string) (*RabbitMQPublisher, error) {
	conn, err := amqp.Dial(amqpURI)
	if err != nil {
		return nil, err
	}

	channel, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	if _, err := channel.QueueDeclare(queueName, true, false, false, false, nil); err != nil {
		return nil, err
	}

	return &RabbitMQPublisher{channel: channel}, nil
}

func (p *RabbitMQPublisher) PublicarPedidoConfirmado(pedidoID, clienteID, productoID string) error {
	evento := PedidoConfirmado{
		Tipo:       "pedido.confirmado",
		PedidoID:   pedidoID,
		ClienteID:  clienteID,
		ProductoID: productoID,
	}

	body, err := json.Marshal(evento)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = p.channel.PublishWithContext(ctx, "", queueName, false, false, amqp.Publishing{
		ContentType:  "application/json",
		DeliveryMode: amqp.Persistent,
		Body:         body,
	})
	if err != nil {
		return err
	}

	log.Printf("[EVENTO] Publicado pedido.confirmado en %q: %s", queueName, body)
	return nil
}
