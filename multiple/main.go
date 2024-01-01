package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Post("/", func(c *fiber.Ctx) error {
		form, err := c.MultipartForm()
		if err != nil {
			return err
		}

		expectedKey := "documents"

		files, keyExists := form.File[expectedKey]
		if !keyExists {
			return fmt.Errorf("Missing expected form key: %s", expectedKey)
		}

		for _, file := range files {
			fmt.Println(file.Filename, file.Size, file.Header["Content-Type"][0])

			err := c.SaveFile(file, fmt.Sprintf("./%s", file.Filename))
			if err != nil {
				return err
			}
		}
		return nil
	})

	log.Fatal(app.Listen(":3000"))
}
