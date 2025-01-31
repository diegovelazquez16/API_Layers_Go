package routes

import (
	"github.com/gin-gonic/gin"
	"holamundo/infraestructure/controllers"
)

func ProductRoutes(
	router *gin.Engine,
	createController *controllers.ProductController,
	getController *controllers.ProductGetController,
	updateController *controllers.ProductUpdateController,

) {
	group := router.Group("/products")
	{
		group.POST("/new", createController.CreateProduct)        
		group.GET("/get/:id", getController.Get)  
		group.PUT("/:id", updateController.Update)

	}
}
