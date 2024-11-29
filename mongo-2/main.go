package main

import (
	config "echo-mongo-api/configs"
	"echo-mongo-api/handlers"
	"fmt"
	"log"

	"github.com/labstack/echo/v4"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	client, err := config.ConnectMongoDB(cfg.MongoURI)
	if err != nil {
		log.Fatal(err)
	}

	db := client.Database("api_test")
	collection := db.Collection("products")

	e := echo.New()

	productHandler := handlers.ProductHandler{Collection: collection}

	e.GET("/", productHandler.GetProducts)
	e.POST("/", productHandler.CreateProduct)

	appPort := fmt.Sprintf(":%s", cfg.AppPort)
	e.Logger.Fatal(e.Start(appPort))
}
