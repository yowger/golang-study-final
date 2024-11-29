package main

import (
	config "echo-mongo-api/configs"
	"log"

	"github.com/labstack/echo/v4"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	e := echo.New()

	e.Logger.Fatal(e.Start(":3000"))

	db, err := config.ConnectMongoDB(cfg.MongoURI)
	if err != nil {
		log.Fatal(err)
	}
}
