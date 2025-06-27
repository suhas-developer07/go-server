package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type envConfig struct {  //innitializing a structure to access variable 
	APP_PORT string
}

func (e *envConfig) LoadConfig() {
	err := godotenv.Load()  //load the env returns error

	if err != nil {
		log.Panic("ENV File not loaded")
	}

	e.APP_PORT = loadstring("APP_PORT", ":8080")
}

var Config envConfig //make that varible first letter capital becouse of to access that in another package 

func init() {  //this function runs first its like a main function in this file 
	Config.LoadConfig()
}

func loadstring(key string, fallback string) string {
	val, ok := os.LookupEnv(key)  //returns string and boolean
	if !ok {
		log.Panic("APP_PORT is not loaded")
		return fallback
	}
	return val

}
