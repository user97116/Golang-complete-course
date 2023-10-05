package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
)

type Todo struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

var todos = []Todo{
	{1, "Buy groceries"},
	{2, "Complete assignment"},
}

func getAllTodos(c *fiber.Ctx) error {
	return c.JSON(todos)
}

func createTodo(c *fiber.Ctx) error {
	var newTodo Todo
	if err := c.BodyParser(&newTodo); err != nil {
		return err
	}

	newTodo.ID = len(todos) + 1
	todos = append(todos, newTodo)

	return c.JSON(newTodo)
}

func main() {
	app := fiber.New()

	app.Get("/todos", getAllTodos)
	app.Post("/todos", createTodo)

	log.Fatal(app.Listen(":8080"))
}
