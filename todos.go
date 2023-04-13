package main

import "github.com/gofiber/fiber/v2"

type Todo struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
	Body  string `json:"body"`
}

var todos []Todo

func health(c *fiber.Ctx) error {
	return c.SendString("Ok!")
}

func addTodo(c *fiber.Ctx) error {
	todo := &Todo{}

	if err := c.BodyParser(todo); err != nil {
		return err
	}

	todo.ID = len(todos) + 1

	todos = append(todos, *todo)

	return c.JSON(todos)
}

func updateTodo(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	if err != nil {
		return c.Status(401).SendString("Invalid id")
	}

	for i, t := range todos {
		if t.ID == id {
			todos[i].Done = true
			break
		}
	}

	return c.JSON(todos)
}

func getTodos(c *fiber.Ctx) error {
	return c.JSON(todos)
}
