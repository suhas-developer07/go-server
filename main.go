package main

import (
	"context"
	"net/http"

	"github.com/suhas-developer07/Golang/config"
	"github.com/suhas-developer07/Golang/db"
	"github.com/suhas-developer07/Golang/routes"
)

func main() {
	handler := routes.Mountroutes()
	db.InitDb() //database calling

	server := &http.Server{
		Addr:    config.Config.APP_PORT,  //Accessing Addr through env variable 
		Handler: handler,
	}
	defer db.DB.Close(context.Background()) //this will runs on unmount means when we off the server 

	server.ListenAndServe()  //listening the server in localhost
}
