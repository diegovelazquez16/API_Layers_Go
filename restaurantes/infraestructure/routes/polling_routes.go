package routes

import (
	"holamundo/users/infraestructure/controllers"
	"github.com/gin-gonic/gin"
)

func PollingRoutes(router *gin.Engine, pollingController *controllers.PollingController){
	router.GET("/polling/short", pollingController.ShortPolling)
	router.GET("/polling/notifications", pollingController.NotificationPolling)
}