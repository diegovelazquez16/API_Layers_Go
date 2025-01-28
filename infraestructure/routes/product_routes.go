package routes

import (
	"github.com/gin-gonic/gin"
	"holamundo/infraestructure/controllers"
)
func ProductRoutes(router *gin.Engine, productController *controllers.ProductController) {
	group := router.Group("/products")
	{
		group.POST("", productController.CreateProduct)

	}
}