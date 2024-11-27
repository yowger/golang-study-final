package main

import (
	"fmt"
	"log"
	"myapp/config"
	"net/http"
)

func main() {
	cfg := config.LoadConfig()

	server := &http.Server{
		Addr:    ":" + cfg.AppPort,
		Handler: nil,
	}

	fmt.Println("Starting server on port", cfg.AppPort)
	if err := server.ListenAndServe(); err != nil {
		log.Fatal("failed to start server: ", err)
	}

}
