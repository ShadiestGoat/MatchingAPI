package main

import (
	// "encoding/json"
	"github.com/gofiber/fiber/v2"
)

const PORT = "3000"

func main() {
	// TODO: Setup db!

	app := fiber.New(fiber.Config{
		AppName: "Shady Dating",
	})

	app.Get("/self/profile", func(c *fiber.Ctx) error {

		// TODO: Send profile info
	})

	app.Listen(":" + PORT)
}
