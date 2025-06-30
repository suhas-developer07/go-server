package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/suhas-developer07/Golang/config"
)

var DB *pgx.Conn

func InitDb() {  
	var err error
	DB, err = pgx.Connect(context.Background(), config.Config.POSTGRES_URL )
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
    

	err = DB.Ping(context.Background())
	if err != nil {
		log.Fatalf("error connecting database")
		os.Exit(1)
	}

 log.Printf("Connected with database")
}