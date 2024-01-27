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
		SigningKey: jwtware.SigningKey{Key: []byte(constants.AuthSignKey)},
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
	ingredient.Put("/:id", handlers.IngredientUpdate)

	// Meal
	meal := api.Group("/meal")
	meal.Get("/", handlers.MealListByUser)
	meal.Get("/all", handlers.MealList)
	meal.Post("/", handlers.MealCreate)
	meal.Get("/:id", handlers.MealRetrieve)
	meal.Put("/:id", handlers.MealUpdate)
	meal.Get("/:id/ingredient", handlers.MealListIngredients)
	meal.Post("/ingredient", handlers.MealAddIngredient)
	meal.Put("/ingredient/:id", handlers.MealIngredientUpdate)
	meal.Delete("/ingredient/:id", handlers.MealIngredientDelete)

	//	Measurement
	measurement := api.Group("/measurement")
	measurement.Get("/", handlers.MeasurementListByUserId)
	measurement.Post("/", handlers.MeasurementCreate)
	measurement.Post("/:id", handlers.MeasurementRetrieve)
	measurement.Put("/:id", handlers.MeasurementUpdate)
	measurement.Delete("/:id", handlers.MeasurementDelete)
}
