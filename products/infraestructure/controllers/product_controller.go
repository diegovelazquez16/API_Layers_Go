package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"holamundo/products/aplication/usecase"
	"holamundo/products/domain/models"
)

type ProductController struct {
	CreateProductUC *aplication.CreateProductUseCase
}

func (c *ProductController) CreateProduct(ctx *gin.Context) {
	var product models.Product
	if err := ctx.ShouldBindJSON(&product); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := c.CreateProductUC.Execute(&product)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Product creado",
		"product": product,
	})
}
