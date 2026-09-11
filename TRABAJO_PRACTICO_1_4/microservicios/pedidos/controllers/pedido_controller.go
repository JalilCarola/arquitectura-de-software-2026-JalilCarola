package controllers

import (
	"errors"
	"net/http"

	"pedidos/services"

	"github.com/gin-gonic/gin"
)

type PedidoController struct {
	service *services.PedidoService
}

func NewPedidoController(service *services.PedidoService) *PedidoController {
	return &PedidoController{service: service}
}

type confirmarPedidoRequest struct {
	ClienteID  string `json:"cliente_id" binding:"required"`
	ProductoID string `json:"producto_id" binding:"required"`
}

func (c *PedidoController) Confirmar(ctx *gin.Context) {
	var req confirmarPedidoRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "cliente_id y producto_id son requeridos"})
		return
	}

	pedido, err := c.service.Confirmar(req.ClienteID, req.ProductoID)
	if err != nil {
		if errors.Is(err, services.ErrProductoInexistente) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "producto inexistente"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "error interno"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"mensaje":   "pedido confirmado",
		"pedido_id": pedido.ID,
	})
}
