package routes

import (
	"holamundo/products/infraestructure/controllers"
	"github.com/gin-gonic/gin"
)

func LPollingRoutes (router *gin.Engine, lPollingController *controllers.LPollingController ){
	router.GET("/polling/long", lPollingController.LongPolling)
}