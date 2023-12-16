package router

import (
	"food-backend/src/handlers"
	"github.com/gofiber/fiber/v2"
)

func placeholderHandler(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"response": ctx.OriginalURL()})
}

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	// Auth
	auth := api.Group("/auth")

	auth.Post("/", handlers.AuthUser)
	auth.Post("/refresh", handlers.RefreshTokens)

	//	User
	user := api.Group("/user")
	user.Get("/", placeholderHandler)
	user.Post("/", placeholderHandler)
	user.Get("/:id", placeholderHandler)
	user.Put("/:id", placeholderHandler)
}
