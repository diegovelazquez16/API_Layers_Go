package controllers

import (
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
	"holamundo/products/aplication/usecase"
)

type ProductGetController struct {
	GetProductUC *aplication.GetProductUseCase
}

func (c *ProductGetController) Get(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	product, err := c.GetProductUC.Execute(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, product)
}

