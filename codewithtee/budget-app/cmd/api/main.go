package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

type Application struct {
	logger echo.Logger
	server *echo.Echo
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading env file")
	}

	appPort := os.Getenv("APP_PORT")
	appAddress := fmt.Sprintf("localhost:%s", appPort)

	e := echo.New()

	// e.Use(middleware.Logger())
	// e.Use(middleware.)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/:id", func(c echo.Context) error {
		id := c.Param("id")
		return c.String(http.StatusOK, id)
	})

	e.Logger.Fatal(e.Start(appAddress))
}
