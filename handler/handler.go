package handler

import (
	"api/database"
	"api/models"
	"api/validators"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

// Hello hanlde api status
func AllActress(c *fiber.Ctx) error {

	users := []models.Actress{}

	database.ReturnDatabase().Find(&users)

	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"data": users,
	})
}

func DeleteActress(c *fiber.Ctx) error {
	act := c.AllParams()
	database.ReturnDatabase().Delete(&models.Actress{}, act["id"])
	return c.SendString("ff")
}

func CreateAllActress(c *fiber.Ctx) error {
	var actress models.Actress

	a := new(validators.Actress)

	if err := c.BodyParser(a); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			fiber.Map{
				"error": "Internal server error",
			},
		)
	}

	u := database.ReturnDatabase().FirstOrCreate(&actress, &models.Actress{Name: a.Name, Age: a.Age})
	if u.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			fiber.Map{
				"error": "Internal server error",
			},
		)
	}

	if u.RowsAffected == 1 {
		return c.Status(fiber.StatusCreated).JSON(
			fiber.Map{
				"message": "User created successfully",
			},
		)
	}

	return c.Status(fiber.StatusBadRequest).JSON(
		fiber.Map{
			"error": "Username already exists",
		},
	)
}
