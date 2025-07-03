package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/suhas-developer07/Golang/db"
)

func SaveTask(ctx *gin.Context) {
	var payload db.PostTaskPayload
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "unable to read the body",
			"error":   err,
		})
		return
	}

	id, err := db.TaskRepository.SaveTaskQuery(payload)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": true,
			"msg":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"error": "false",
		"id":    id,
	})
}

func ReadTask(ctx *gin.Context) {

	tasks, err := db.TaskRepository.ReadTaskQuery()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "true",
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"error": false,
		"data":  tasks,
	})

}

func DeletTask(ctx *gin.Context) {
	var payload db.DelteTaskType

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": true,
			"msg":   err.Error(),
		})
		return
	}

	err := db.TaskRepository.DeleteTaskQuery(payload.ID)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": true,
			"msg":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"error": false,
		"msg":   "Task deleted successfully",
	})

}

func UpdateTask(ctx *gin.Context) {
	var payload db.UpdateTaskType

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": true,
			"msg":   err.Error(),
		})
		return
	}

	task, err := db.TaskRepository.GetTaskById(payload.ID)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": true,
			"msg":   err.Error(),
		})
		return
	}

	if payload.Title == "" {
		payload.Title = task.Title
	}
	if payload.Description == "" {
		payload.Description = task.Description
	}
	if payload.Status == "" {
		payload.Status = task.Status
	}

	UpdateTaskErr := db.TaskRepository.UpdateTask(payload)

	if UpdateTaskErr != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": true,
			"msg":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"error": false,
		"data":  payload,
	})
}
