package controllers

import (
	"net/http"
	"time"
	"holamundo/users/aplication/usecase"
	"github.com/gin-gonic/gin"
)

type PollingController struct {
	GetAllUsersUC *usecase.GetAllUsersUseCase
}

func (c *PollingController) ShortPolling(ctx *gin.Context){
	users, err := c.GetAllUsersUC.Execute()
	if err != nil{
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
		ctx.JSON(http.StatusOK, users)
}

func (c *PollingController) NotificationPolling(ctx *gin.Context){
	time.Sleep(5 * time.Second)
	notifications := [] string{
		"Esta es una notificación",
	}
	ctx.JSON(http.StatusOK, gin.H{"notifications": notifications})
}