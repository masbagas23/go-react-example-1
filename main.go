package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

type Todo struct {
	ID        int    `json:"id"`
	Completed bool   `json:"completed"`
	Body      string `json:"body"`
}

func main() {
	app := fiber.New()

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	PORT := os.Getenv("PORT")

	todos := []Todo{}

	app.Get("/api/todos", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(todos)
	})

	// Create todo
	app.Post("/api/todos", func(c *fiber.Ctx) error {
		todo := &Todo{}

		if err := c.BodyParser(todo); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": err.Error()})
		}

		if todo.Body == "" {
			return c.Status(400).JSON(fiber.Map{"error": "body is required"})
		}

		todo.ID = len(todos) + 1

		todos = append(todos, *todo)

		return c.Status(201).JSON(todo)
	})

	// Update todo
	app.Patch("/api/todos/:id", func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")

		if err != nil {
			return c.Status(400).JSON(fiber.Map{"error": err.Error()})
		}

		for _, todo := range todos {
			if todo.ID == id {
				todo.Completed = true
				return c.Status(200).JSON(todo)
			}
		}

		return c.Status(404).JSON(fiber.Map{"error": "todo not found"})
	})

	// Delete todo
	app.Delete("/api/todos/:id", func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")

		if err != nil {
			return c.Status(400).JSON(fiber.Map{"error": err.Error()})
		}

		for i, todo := range todos {
			if todo.ID == id {
				todos = append(todos[:i], todos[i+1:]...)
				return c.Status(200).JSON(fiber.Map{"message": "todo deleted"})
			}
		}
		return c.Status(404).JSON(fiber.Map{"error": "todo not found"})
	})

	log.Fatal(app.Listen(":" + PORT))
}
