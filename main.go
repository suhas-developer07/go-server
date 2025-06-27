package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/suhas-developer07/Golang/config"
	"github.com/suhas-developer07/Golang/db"
)

func main() {
	handler := gin.Default()  //gin framework that used to create http server in golang
	handler.GET("/ping", func(ctx *gin.Context) { //route handler
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	db.InitDb() //database calling

	server := &http.Server{
		Addr:    config.Config.APP_PORT,  //Accessing Addr through env variable 
		Handler: handler,
	}

	server.ListenAndServe()  //listening the server in localhost
}
