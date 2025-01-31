package routes

import (
	"github.com/gin-gonic/gin"
	"holamundo/infraestructure/controllers"
)

func ProductRoutes(
	router *gin.Engine,
	createController *controllers.ProductController,
	getController *controllers.ProductGetController,

) {
	group := router.Group("/products")
	{
		group.POST("", createController.CreateProduct)        // POST /products
		group.GET("/:id", getController.Get)           // GET /products/:id
	}
}
