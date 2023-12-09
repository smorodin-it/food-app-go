package handlers

import (
	"fmt"
	"food-backend/src/forms"
	"food-backend/src/repositories"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"time"
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

	// Check user password

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(formAuth.Password))
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}

	// Generate tokens

	token := jwt.New(jwt.SigningMethodHS512)

	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = user.Id
	claims["email"] = user.Email

	exp := time.Now().Add(time.Minute)

	claims["exp"] = exp.Unix()

	t, err := token.SignedString([]byte("3be9067a-0869-4291-ae46-1e943f05642a"))
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}

	// Return tokens

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"token":        t,
		"refreshToken": uuid.New().String(),
	})
}
