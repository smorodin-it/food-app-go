package handlers

import (
	"food-backend/src/forms"
	"food-backend/src/repositories"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"time"
)

func AuthUser(ctx *fiber.Ctx) error {
	formAuth := new(forms.FormAuth)
	err := ctx.BodyParser(formAuth)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err})
	}

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
	claims["exp"] = time.Now().Add(time.Minute).Unix()

	at, err := token.SignedString([]byte("3be9067a-0869-4291-ae46-1e943f05642a"))
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}

	rt, err := ur.SetRefreshToken(user.Id)

	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}

	// Return tokens

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"token":        at,
		"refreshToken": rt,
	})
}

func RefreshTokens (ctx *fiber.Ctx) error {
	return nil
}