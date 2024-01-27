package handlers

import (
	"food-backend/src/forms"
	"food-backend/src/services"
	"food-backend/src/utils"
	"github.com/gofiber/fiber/v2"
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
	f := new(forms.FormAuth)
	err := ctx.BodyParser(f)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err})
	}

	as := new(services.AuthService)
	tokens, err := as.AuthUser(*f)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err})
	}

	utils.SetTokensToResponse(ctx, *tokens)

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

	as := new(services.AuthService)
	tokens, err := as.RefreshTokens(*form)

	utils.SetTokensToResponse(ctx, *tokens)

	return ctx.SendStatus(fiber.StatusOK)
}
