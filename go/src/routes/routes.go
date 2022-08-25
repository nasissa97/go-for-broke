package routes

import (
	"api/src/database"

	"github.com/gofiber/fiber/v2"
)

func GetPeople(c *fiber.Ctx) error {
	people := database.GetData()
	return c.Status(200).JSON(people)
}
