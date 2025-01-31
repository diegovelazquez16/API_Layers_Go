package controllers

import (
	"holamundo/products/aplication/usecase"
	"holamundo/products/domain/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)
type ProductUpdateController struct{
	UpdateProductUC *aplication.UpdateProductUsecase
}

func (c *ProductUpdateController) Update(ctx *gin.Context){
	var product models.Product

	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	
	if err != nil{
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID invalido"})
		return
	}

	if err := ctx.ShouldBindJSON(&product); err != nil{
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	product.ID = uint(id)

		err = c.UpdateProductUC.Execute(&product)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}

		ctx.JSON(http.StatusOK, gin.H{"message": "Product actualizado exitosamente"})
}
