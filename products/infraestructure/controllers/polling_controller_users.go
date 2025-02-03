package controllers

import (
	"holamundo/products/aplication/usecase"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type LPollingController struct{
	GetAllProductsUseCase *aplication.GetAllProductsUseCase
}

func (c *LPollingController) LongPolling(ctx *gin.Context){
	ctx.Writer.Header().Set("Conten-Type", "aplication/json")
	ctx.Writer.Header().Set("Transfer-Encoding", "chunked")
	ctx.Writer.Flush() //Si no es relevante eliminarlo

	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()
		timeout := time.After(30 * time.Second)

		for {
			select{
			case <- timeout:
				ctx.JSON(http.StatusGatewayTimeout, gin.H{"message": "No hay actualizaciones"})
				return
			case <- ticker.C:
				products, err := c.GetAllProductsUseCase.Execute()
				if err != nil{
					ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
					return
				}
				if len(products) > 0{
					ctx.SSEvent("data", products)
					ctx.Writer.Flush()
				}

			}
		}
}