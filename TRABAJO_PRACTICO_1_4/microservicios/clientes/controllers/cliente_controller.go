package controllers

import (
	"errors"
	"net/http"

	"clientes/repositories"
	"clientes/services"

	"github.com/gin-gonic/gin"
)

type ClienteController struct {
	service *services.ClienteService
}

func NewClienteController(service *services.ClienteService) *ClienteController {
	return &ClienteController{service: service}
}

type crearClienteRequest struct {
	Nombre string `json:"nombre" binding:"required"`
	Email  string `json:"email" binding:"required"`
}

func (c *ClienteController) Crear(ctx *gin.Context) {
	var req crearClienteRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "nombre y email son requeridos"})
		return
	}

	cliente := c.service.Crear(req.Nombre, req.Email)
	ctx.JSON(http.StatusCreated, cliente)
}

func (c *ClienteController) ObtenerPorID(ctx *gin.Context) {
	cliente, err := c.service.ObtenerPorID(ctx.Param("id"))
	if err != nil {
		if errors.Is(err, repositories.ErrClienteNoEncontrado) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "cliente no encontrado"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "error interno"})
		return
	}

	ctx.JSON(http.StatusOK, cliente)
}
