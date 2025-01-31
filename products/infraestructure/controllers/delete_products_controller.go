package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"holamundo/products/aplication/usecase"
)

type ProductDeleteController struct {
	DeleteProductUC *aplication.DeleteProductUseCase
}

func (c *ProductDeleteController) Delete(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	err = c.DeleteProductUC.Execute(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Producto eliminado exitosamente"})
}
