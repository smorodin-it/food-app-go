package router

import (
	"food-backend/src/constants"
	"food-backend/src/database"
	"food-backend/src/handlers"
	"food-backend/src/repositories"
	"food-backend/src/services"
	"food-backend/src/utils"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"log"
)

func SetupRoutes(app *fiber.App) {
	// Initialize database
	db, err := database.Connect(database.DbConfig)
	if err != nil {
		log.Fatal(err)
	}

	// Initialize repositories and services needed for public routes
	userRepo := repositories.NewUserRepository(db)
	authServ := services.NewAuthService(userRepo)

	api := app.Group("/api")

	// Auth
	auth := api.Group("/auth")
	auth.Post("/", handlers.AuthUser(authServ))
	auth.Post("/refresh", handlers.RefreshTokens(authServ))

	api.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(constants.AuthSignKey)},
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			return utils.GetResponseError(ctx, fiber.StatusUnauthorized, err)
		},
	},
	))

	// Initialize private repositories
	ingredientRepo := repositories.NewIngredientRepo(db)
	mealRepo := repositories.NewMealRepository(db)
	measurementRepo := repositories.NewMeasurementRepository(db)
	inventoryRepo := repositories.NewInventoryRepo(db)

	// Initialize private services
	//userServ - TODO: Implement
	ingredientServ := services.NewIngredientService(ingredientRepo)
	mealServ := services.NewMealService(mealRepo)
	measurementServ := services.NewMeasurementService(measurementRepo)
	inventoryServ := services.NewInventoryService(inventoryRepo)

	// Private handlers

	//	User
	//user := api.Group("/user")
	//user.Get("/", placeholderHandler)
	//user.Post("/", placeholderHandler)
	//user.Get("/:id", placeholderHandler)
	//user.Put("/:id", placeholderHandler)

	// Ingredient
	ingredient := api.Group("/ingredient")
	ingredient.Get("/", handlers.IngredientList(ingredientServ))
	ingredient.Post("/", handlers.IngredientCreate(ingredientServ, authServ))
	ingredient.Get("/:id", handlers.IngredientRetrieve(ingredientServ))
	ingredient.Put("/:id", handlers.IngredientUpdate(ingredientServ))

	// Meal
	meal := api.Group("/meal")
	meal.Get("/", handlers.MealListByUser(mealServ))
	meal.Get("/all", handlers.MealList(mealServ))
	meal.Post("/", handlers.MealCreate(mealServ, authServ))
	meal.Get("/:id", handlers.MealRetrieve(mealServ))
	meal.Put("/:id", handlers.MealUpdate(mealServ))
	meal.Get("/:id/ingredient", handlers.MealListIngredients(mealServ))
	meal.Post("/ingredient", handlers.MealAddIngredient(mealServ))
	meal.Put("/ingredient/:id", handlers.MealIngredientUpdate(mealServ))
	meal.Delete("/ingredient/:id", handlers.MealIngredientDelete(mealServ))

	//	Measurement
	measurement := api.Group("/measurement")
	measurement.Get("/", handlers.MeasurementListByUserId(measurementServ))
	measurement.Post("/", handlers.MeasurementCreate(measurementServ))
	measurement.Post("/:id", handlers.MeasurementRetrieve(measurementServ))
	measurement.Put("/:id", handlers.MeasurementUpdate(measurementServ))
	measurement.Delete("/:id", handlers.MeasurementDelete(measurementServ))

	//	Inventory
	inventory := api.Group("/inventory")
	inventory.Get("/", handlers.InventoryList(inventoryServ))
	inventory.Post("/", handlers.InventoryCreate(inventoryServ, authServ))
	inventory.Get("/:id", handlers.InventoryRetrieve(inventoryServ))
	inventory.Put("/:id", handlers.InventoryUpdate(inventoryServ))
}
