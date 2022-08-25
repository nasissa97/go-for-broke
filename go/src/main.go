package main

import (
	"api/src/routes"

	"github.com/gofiber/fiber/v2"
)

func setUpRoutes(app *fiber.App) {
	app.Get("/people", routes.GetPeople)
	// app.Post("/people", routes.AddPerson)
	// app.Delete("/people", routes.DeletePerson)

}

func main() {
	app := fiber.New()

	setUpRoutes(app)

	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404)
	})

	app.Listen(":3000")
}
