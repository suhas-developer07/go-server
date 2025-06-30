package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/suhas-developer07/Golang/routes/handlers"
)


func Mountroutes() *gin.Engine {
handler := gin.Default()

	handler.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Ok from gin",
		})
	})

	handler.POST("/task", handlers.SaveTask)

	return  handler
}