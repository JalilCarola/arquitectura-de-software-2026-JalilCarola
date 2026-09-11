package main

import (
	"log"

	"clientes/controllers"
	"clientes/repositories"
	"clientes/services"

	"github.com/gin-gonic/gin"
)

func main() {
	repository := repositories.NewClienteRepository()
	service := services.NewClienteService(repository)
	controller := controllers.NewClienteController(service)

	router := gin.Default()
	router.POST("/clientes", controller.Crear)
	router.GET("/clientes/:id", controller.ObtenerPorID)

	log.Println("Microservicio de clientes escuchando en http://localhost:8081")
	if err := router.Run(":8081"); err != nil {
		log.Fatal(err)
	}
}
