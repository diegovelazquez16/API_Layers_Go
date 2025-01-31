package routes

import (
	
	"holamundo/users/infraestructure/controllers"
	"github.com/gin-gonic/gin"

)

func UserRoutes(
	router *gin.Engine,
	createController *controllers.UserCreateController,
	getAllController *controllers.UserGetAllController,

) {
	group := router.Group("/users")
	{
		group.POST("", createController.Create)          
		group.GET("", getAllController.GetAll)           
    
	}
}