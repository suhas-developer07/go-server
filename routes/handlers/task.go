package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/suhas-developer07/Golang/db"
)

func SaveTask(ctx *gin.Context){
		var payload db.PostTaskPayload
		if err := ctx.ShouldBindJSON(&payload);

		 err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "unable to read the body",
				"error":   err,
			})
			return
		}
	   
		id, err := db.TaskRepository.SaveTaskQuery(payload)

		if err !=nil {
			ctx.JSON(http.StatusInternalServerError,gin.H{
				"error":true,
				"msg":err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"error": "false",
			"id":id,
		})
}

func ReadTask(ctx *gin.Context){
   
	tasks,err := db.TaskRepository.ReadTaskQuery()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError,gin.H{
			"error":"true",
			"message":err.Error(),
		})
		return 
	}
	ctx.JSON(http.StatusOK, gin.H{
		"error":false,
		"data":tasks,
	})
		
}

