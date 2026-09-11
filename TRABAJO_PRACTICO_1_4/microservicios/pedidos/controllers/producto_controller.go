package controllers

import (
	"net/http"

	"pedidos/services"

	"github.com/gin-gonic/gin"
)

type ProductoController struct {
	service *services.ProductoService
}

func NewProductoController(service *services.ProductoService) *ProductoController {
	return &ProductoController{service: service}
}

func (c *ProductoController) ListarTodos(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"productos": c.service.ObtenerTodos()})
}
