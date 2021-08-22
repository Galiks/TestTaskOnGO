package main

import (
	"GoProject/storage"
	"log"

	"github.com/gofiber/fiber/v2"
)

var App *fiber.App = fiber.New()

func main() {
	App.Get("/json/hackers", func(c *fiber.Ctx) error {
		c.Accepts("application/json")
		hackers, err := storage.GetValues()
		if err != nil {
			return c.JSON(&fiber.Map{
				"success": false,
				"data":    nil,
				"error":   err.Error(),
				"status":  "500",
			})
		}
		return c.JSON(&fiber.Map{
			"success": true,
			"data":    hackers,
			"error":   nil,
			"status":  "200",
		})
	})
	App.Get("/", func(c *fiber.Ctx) error {
		return c.Redirect("/json/hackers")
	})

	log.Fatal(App.Listen(":8010"))
}
