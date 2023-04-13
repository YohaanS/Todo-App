package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:5173",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	app.Get("api/healthcheck", health)

	app.Post("api/todos", addTodo)

	app.Patch("api/todos/:id/done", updateTodo)

	app.Get("api/todos", getTodos)

	log.Fatal(app.Listen(":4000"))
}
