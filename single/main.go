package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Post("/", func(c *fiber.Ctx) error {
		file, err := c.FormFile("document")
		if err != nil {
			return err
		}

		return c.SaveFile(file, fmt.Sprintf("./%s", file.Filename))
	})

	log.Fatal(app.Listen(":3000"))
}
