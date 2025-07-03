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
  //grouping the task routes
	TaskRoutes := handler.Group("/task")   
	{
		TaskRoutes.POST("/", handlers.SaveTask)
	    TaskRoutes.GET("/",handlers.ReadTask)
		TaskRoutes.DELETE("/",handlers.DeletTask)
	    TaskRoutes.PATCH("/",handlers.UpdateTask)

	}
 //handles the no route api
	handler.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound,gin.H{
			"message":"Route not found",
		})
	})

	return  handler
}