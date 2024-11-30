package main

import (
	"fmt"
	"net/http"
	"strconv"
	"sync"

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
	users = map[int]*user{}
	seq   = 1
	lock  = sync.Mutex{}
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
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, users[id])
}

func getAllUsers(c echo.Context) error {
	lock.Lock()
	defer lock.Unlock()

	return c.JSON(http.StatusOK, users)
}

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/", createUser)
	e.GET("/:id", getUserByID)
	e.GET("/", getAllUsers)

	e.Logger.Fatal(e.Start(":8080"))
}
