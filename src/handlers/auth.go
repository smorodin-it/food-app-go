package handlers

import (
	"food-backend/src/forms"
	"food-backend/src/repositories"
	"food-backend/src/services"
	"food-backend/src/utils"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

// AuthUser is a function to get tokens to work with app
// @Summary Get Auth tokens using credentials
// @Description Get access and refresh tokens
// @Tags Auth
// @Accept json
// @Param credentials body forms.FormAuth true "Auth user"
// @Produce plain
// @Success 200
// @Router /auth [post]
func AuthUser(ctx *fiber.Ctx) error {
	formAuth := new(forms.FormAuth)
	err := ctx.BodyParser(formAuth)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err})
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

	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}

	// Generate tokens

	as := new(services.AuthService)

	t, err := as.GenerateTokens(user)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	err = ur.SetRefreshToken(user.Id, t.RefreshToken)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	utils.SetTokensToCookies(ctx, *t)

	return ctx.SendStatus(fiber.StatusOK)
}

// RefreshTokens is a function to refresh tokens
// @Summary Get new tokens using refresh token
// @Description Get access and refresh tokens
// @Tags Auth
// @Produce plain
// @Success 200
// @Router /auth/refresh [post]
func RefreshTokens(ctx *fiber.Ctx) error {
	// Get refresh token from cookies
	form := new(forms.RefreshForm)

	err := ctx.CookieParser(form)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	err = form.Validate()
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	// Get user from db by refresh token

	ur := new(repositories.UserRepository)

	user, err := ur.GetByRefreshToken(form.RefreshToken)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	// Generate and update tokens
	as := new(services.AuthService)

	t, err := as.GenerateTokens(user)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	err = ur.SetRefreshToken(user.Id, t.RefreshToken)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	utils.SetTokensToCookies(ctx, *t)

	return ctx.SendStatus(fiber.StatusOK)
}
