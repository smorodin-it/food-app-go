package handlers

import (
	"fmt"
	"food-backend/src/forms"
	"food-backend/src/repositories"
	"github.com/gofiber/fiber/v2"
)

func HandleAuth(ctx *fiber.Ctx) error {
	formAuth := new(forms.FormAuth)
	err := ctx.BodyParser(formAuth)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err})
	}

	fmt.Println(formAuth)

	ur := new(repositories.UserRepository)

	//	Get user by email

	user, err := ur.GetByEmail(formAuth.Email)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.Status(fiber.StatusOK).JSON(user)

	// Check user password

	// Generate tokens

	// Return tokens

	return nil
}
