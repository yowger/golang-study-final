package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type (
	user struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}
)

var (
	users          = map[int]*user{}
	seq            = 1
	lock           = sync.Mutex{}
	allowedOrigins = []string{"https://labstack.com", "https://labstack.net"}
)

func createUser(c echo.Context) error {
	lock.Lock()
	defer lock.Unlock()

	fmt.Println("create user")

	u := &user{
		ID: seq,
	}
	if err := c.Bind(u); err != nil {
		return err
	}

	users[u.ID] = u
	seq++

	fmt.Println("users: ", users)

	return c.JSON(http.StatusOK, u)
}

func getUserByID(c echo.Context) error {
	lock.Lock()
	defer lock.Unlock()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println("Invalid user ID error: ", err)

		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid user ID"})
	}

	user, exists := users[id]
	if !exists {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "User not found"})
	}

	return c.JSON(http.StatusOK, user)
}

func getAllUsers(c echo.Context) error {
	lock.Lock()
	defer lock.Unlock()

	return c.JSON(http.StatusOK, users)
}

func updateUserByID(c echo.Context) error {
	lock.Lock()
	defer lock.Unlock()

	u := new(user)
	if err := c.Bind(u); err != nil {
		return err
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println("Invalid user ID error: ", err)

		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid user ID"})
	}

	users[id].Name = u.Name

	return c.JSON(http.StatusOK, users[id])
}

func deleteUserByID(c echo.Context) error {
	lock.Lock()
	defer lock.Unlock()

	id, _ := strconv.Atoi(c.Param("id"))

	delete(users, id)

	return c.NoContent(http.StatusNoContent)
}

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: allowedOrigins,
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

	e.GET("/users", getAllUsers)
	e.POST("/users", createUser)
	e.GET("/users/:id", getUserByID)
	e.PUT("/users/:id", updateUserByID)
	e.DELETE("/users/:id", deleteUserByID)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	go func() {
		err := e.Start(":8080")
		if err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server", err)
		}
	}()

	fmt.Println("START THIS SERVER BABY!")

	// keeps program alive unless terminate signal
	<-ctx.Done()

	fmt.Println("Server shutdown gracefully")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}

}
