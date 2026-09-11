package main

import (
	"log"

	"pedidos/controllers"
	"pedidos/messaging"
	"pedidos/repositories"
	"pedidos/services"

	"github.com/gin-gonic/gin"
)

const amqpURI = "amqp://user:pass@localhost:5672"

func main() {
	publisher, err := messaging.NewRabbitMQPublisher(amqpURI)
	if err != nil {
		log.Fatalf("no se pudo conectar con RabbitMQ: %v", err)
	}

	productoRepo := repositories.NewProductoRepository()
	pedidoRepo := repositories.NewPedidoRepository()

	productoService := services.NewProductoService(productoRepo)
	pedidoService := services.NewPedidoService(pedidoRepo, productoService, publisher)

	productoController := controllers.NewProductoController(productoService)
	pedidoController := controllers.NewPedidoController(pedidoService)

	router := gin.Default()
	router.GET("/productos", productoController.ListarTodos)
	router.POST("/pedidos", pedidoController.Confirmar)

	log.Println("Microservicio de pedidos escuchando en http://localhost:8082")
	if err := router.Run(":8082"); err != nil {
		log.Fatal(err)
	}
}
