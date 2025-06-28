package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
)

var DB *pgx.Conn

func InitDb() {
	urlExample := "postgres://suhas:secrete123@localhost:5432/mydb"  //postgress url 
	var err error
	DB, err = pgx.Connect(context.Background(), urlExample)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer DB.Close(context.Background())  //this will runs on unmount means when we off the server 

	err = DB.Ping(context.Background())
	if err != nil {
		log.Fatalf("error connecting database")
		os.Exit(1)
	}

 log.Printf("Connected with database")
}