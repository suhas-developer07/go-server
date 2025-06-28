package main

import (
	"context"
	"net/http"

	"github.com/suhas-developer07/Golang/config"
	"github.com/suhas-developer07/Golang/db"
	"github.com/suhas-developer07/Golang/routes/handlers"
)

func main() {
	handler := handlers.Mountroutes()
	db.InitDb() //database calling

	server := &http.Server{
		Addr:    config.Config.APP_PORT,  //Accessing Addr through env variable 
		Handler: handler,
	}
	defer db.DB.Close(context.Background()) 

	server.ListenAndServe()  //listening the server in localhost
}
