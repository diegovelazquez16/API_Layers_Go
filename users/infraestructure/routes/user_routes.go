package routes

import (
	
	"holamundo/users/infraestructure/controllers"
	"github.com/gin-gonic/gin"

)

func UserRoutes(router *gin.Engine, createController *controllers.UserCreateController,) {
	group := router.Group("/users")
	{
		group.POST("", createController.Create)     // Crear Usuario
	}
}