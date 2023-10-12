package pkg

func FiberSpa() string {
	return `
package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

const port = ":8887"
const path = "ui"

func main(){
	app := fiber.New()

	app.Static("/", path)

	app.Get("/hi", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	app.Use(cors.New())
	app.Use(logger.New())
	app.Use(recover.New())

	if err := app.Listen(port); err != nil {
		log.Fatal(err)
	}
}
`
}
