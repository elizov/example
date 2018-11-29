package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"app/entity"

	"github.com/joho/godotenv"
	"github.com/mholt/binding"
)

func init() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("main.init -> godotenv.Load", err.Error())
	}
}

func handler(resp http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		user := new(entity.User)
		if err := binding.Bind(req, user); err != nil {
			http.Error(resp, err.Error(), http.StatusBadRequest)
			return
		}

		fmt.Fprintf(resp, "ID:    %d\n", user.ID)
		fmt.Fprintf(resp, "Email: %s\n", user.Email)
		fmt.Fprintf(resp, "Name: %s\n", user.Name)
	}
}

func main() {
	http.HandleFunc("/", handler)

	appUrl := os.Getenv("APP_URL")
	appPort := ":" + os.Getenv("APP_PORT")
	log.Println("The server was run at " + appUrl + appPort)
	if err := http.ListenAndServe(appPort, nil); err != nil {
		log.Fatal("main.main -> http.ListenAndServe", err)
	}
}
