package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PostTaskPayload struct {
	Title       string `json:"Title" binding:"required`
	Description string `json:"Description binding:"required"`
	Status      string `json:"status" binding:"required"`
}

func Mountroutes() *gin.Engine {
	handler := gin.Default()

	handler.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Ok from gin",
		})
	})

	handler.POST("/task", func(ctx *gin.Context) {
		var payload PostTaskPayload
		err := ctx.ShouldBindJSON(&payload)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "unable to read the body",
				"error":   err,
			})
			return
		}
		log.Printf(payload.Status)
		log.Println(payload.Title)

		ctx.JSON(http.StatusOK, gin.H{
			"error": "false",
		})
	})
	return handler
}
