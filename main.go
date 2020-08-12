package main

import (
	"fmt"
	"github.com/yigitsadic/onetimecode/handlers"
	"github.com/yigitsadic/onetimecode/models"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	expiration := 60
	codeStore := models.NewCodeStore(expiration)

	log.Printf("Server is up and running on PORT %s", port)

	http.HandleFunc("/create", handlers.HandleCreate(codeStore))
	http.HandleFunc("/read", handlers.HandleRead(codeStore))

	err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
	if err != nil {
		log.Fatalf("Unable to continue serving cause of %s", err)
	}
}
