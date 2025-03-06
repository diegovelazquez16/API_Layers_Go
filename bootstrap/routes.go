package bootstrap

import "github.com/gin-gonic/gin"

func RegisterRoutes(router *gin.Engine) {
	RegisterProductModule(router)
	RegisterUserModule(router)
}
