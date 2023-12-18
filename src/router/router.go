package router

import (
	"food-backend/src/constants"
	"food-backend/src/handlers"
	jwtware "github.com/gofiber/contrib/jwt"
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

	api.Use(jwtware.New(jwtware.Config{
		SigningKey:  jwtware.SigningKey{Key: []byte(constants.AuthSignKey)},
		TokenLookup: "cookie:token",
	},
	))

	//	User
	//user := api.Group("/user")
	//user.Get("/", placeholderHandler)
	//user.Post("/", placeholderHandler)
	//user.Get("/:id", placeholderHandler)
	//user.Put("/:id", placeholderHandler)

	// Ingredient
	ingredient := api.Group("/ingredient")
	ingredient.Get("/", handlers.IngredientList)
	ingredient.Post("/", handlers.IngredientCreate)
	ingredient.Get("/:id", handlers.IngredientRetrieve)
	ingredient.Put("/:id", placeholderHandler)
}
