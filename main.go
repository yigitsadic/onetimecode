package main

import (
	"context"
	"fmt"
	"github.com/yigitsadic/onetimecode/handlers"
	"github.com/yigitsadic/onetimecode/shared"
	"log"
	"net/http"
	"os"
)

var ctx = context.Background()

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server is up and running on PORT %s", port)

	redisService := shared.NewRedisService()

	http.HandleFunc("/create", handlers.HandleCreate(redisService, &ctx))
	http.HandleFunc("/read", handlers.HandleRead(redisService, &ctx))

	err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
	if err != nil {
		log.Fatalf("Unable to continue serving cause of %s", err)
	}
}
