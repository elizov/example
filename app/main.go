package main

import (
	"log"
	"net/http"
	"os"

	"app/handlers"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("main.init -> godotenv.Load", err.Error())
	}
}

func main() {
	appUrl := os.Getenv("APP_URL")
	appPort := ":" + os.Getenv("APP_PORT")

	log.Println("The server was run at " + appUrl + appPort)

	http.HandleFunc("/", handlers.Handler)
	if err := http.ListenAndServe(appPort, nil); err != nil {
		log.Fatal("main.main -> http.ListenAndServe", err)
	}
}
