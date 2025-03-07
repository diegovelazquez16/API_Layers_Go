package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"holamundo/pedidos/aplication/usecase"
	"holamundo/pedidos/domain/models"
)

type PedidoCreateController struct {
	CreatePedidoUC *usecase.CreatePedidoUseCase
}

func (c *PedidoCreateController) Create(ctx *gin.Context) {
	var Pedido models.Pedido
	if err := ctx.ShouldBindJSON(&Pedido); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := c.CreatePedidoUC.Execute(&Pedido)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Pedido creado de forma exitosa",
		"Pedido": Pedido,
	})
}
