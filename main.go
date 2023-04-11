package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

type Todo struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
	Body  int    `json:"body"`
}

func main() {
	fmt.Println("Hello World")

	app := fiber.New()

	app.Get("/healthcheck", func (c *fiber.Ctx) error {
	return c.SendString("Ok!")
	})

	log.Fatal(app.Listen(":3000"))
}
