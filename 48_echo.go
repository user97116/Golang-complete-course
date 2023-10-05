package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var users = []User{
	{1, "Alice"},
	{2, "Bob"},
}

func getUser(c echo.Context) error {
	id := c.Param("id")
	for _, user := range users {
		if id == fmt.Sprintf("%d", user.ID) {
			return c.JSON(http.StatusOK, user)
		}
	}
	return c.NoContent(http.StatusNotFound)
}

func main() {
	e := echo.New()

	e.GET("/api/users/:id", getUser)

	e.Start(":8080")
}
