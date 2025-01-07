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
	"github.com/gofiber/fiber/v2/middleware/logger"
	"log"
)

func SetupRoutes(app *fiber.App) {
	// Initialize database
	db, err := database.Connect(database.DbConfig)
	if err != nil {
		log.Fatal(err)
	}

	app.Use(logger.New())

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
	ih := handlers.NewIngredientHandler(authServ, ingredientServ)
	ingredient := api.Group("/ingredient")
	ingredient.Get("/", ih.List())
	ingredient.Post("/", ih.Create())
	ingredient.Get("/:id", ih.Retrieve())
	ingredient.Put("/:id", ih.Update())

	// Meal
	mh := handlers.NewMealHandler(authServ, mealServ)
	meal := api.Group("/meal")
	meal.Get("/", mh.ListByUser())
	meal.Get("/all", mh.List())
	meal.Post("/", mh.Create())
	meal.Get("/:id", mh.Retrieve())
	meal.Put("/:id", mh.Update())
	meal.Get("/:id/ingredient", mh.ListIngredient())
	meal.Post("/ingredient", mh.AddIngredient())
	meal.Put("/ingredient/:id", mh.UpdateIngredient())
	meal.Delete("/ingredient/:id", mh.DeleteIngredient())

	//	Measurement
	meh := handlers.NewMeasurementHandler(authServ, measurementServ)
	measurement := api.Group("/measurement")
	measurement.Get("/", meh.ListByUserId())
	measurement.Post("/", meh.Create())
	measurement.Post("/:id", meh.Retrieve())
	measurement.Put("/:id", meh.Update())
	measurement.Delete("/:id", meh.Delete())

	//	Inventory
	inventory := api.Group("/inventory")
	inventory.Get("/", handlers.InventoryList(inventoryServ))
	inventory.Post("/", handlers.InventoryCreate(inventoryServ, authServ))
	inventory.Get("/:id", handlers.InventoryRetrieve(inventoryServ))
	inventory.Put("/:id", handlers.InventoryUpdate(inventoryServ))
}
