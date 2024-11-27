package main

import (
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/:id", func(c echo.Context) error {
		id := c.Param("id")
		return c.String(http.StatusOK, id)
	})
	e.Logger.Fatal(e.Start(":3000"))
}
