package routes

import (
	"github.com/gin-gonic/gin"
	"holamundo/products/infraestructure/controllers"
)

func ProductRoutes(
	router *gin.Engine,
	createController *controllers.ProductController,
	getController *controllers.ProductGetController,
	updateController *controllers.ProductUpdateController,
	deleteController *controllers.ProductDeleteController,
	getAllController *controllers.ProductGetAllController,

) {
	group := router.Group("/products")
	{
		group.POST("/", createController.CreateProduct)        
		group.GET("/:id", getController.Get)  
		group.PUT("/:id", updateController.Update)
		group.DELETE("/:id", deleteController.Delete)
		group.GET("/", getAllController.GetAll)

	}
}
