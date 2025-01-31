package controllers

import (
	"holamundo/aplication/usecase"
	"net/http"
	"github.com/gin-gonic/gin"


)

type ProductGetAllController struct {
	GetAllProductsUC *aplication.GetAllProductsUseCase
}

func (c *ProductGetAllController) GetAll(ctx *gin.Context) {
	products, err := c.GetAllProductsUC.Execute()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, products)
}