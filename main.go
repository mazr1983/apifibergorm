package main

import (
	"github.com/mazr1983/apifibergorm/pkg/connector"
	"github.com/mazr1983/apifibergorm/pkg/route"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	db := connector.Connect()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(map[string]string{
			"message": "Figo API",
		})
	})

	route.Expenses(app, db)
	app.Listen(":3000")
}
