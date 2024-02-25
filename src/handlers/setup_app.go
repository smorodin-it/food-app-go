package handlers

import (
	"food-backend/src/constants"
	"food-backend/src/database"
	"food-backend/src/repositories"
	"food-backend/src/services"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"log"
)

func SetupApp(app *fiber.App) {
	// Initialize database
	db, err := database.Connect(database.DbConfig)
	if err != nil {
		log.Fatal(err.Error())
	}

	// Initialize repositories and services needed for public routes
	userRepo := repositories.NewUserRepository(db)
	authServ := services.NewAuthService(userRepo)

	api := app.Group("/api")

	// Auth
	auth := api.Group("/auth")
	auth.Post("/", AuthUser(authServ))
	auth.Post("/refresh", RefreshTokens(authServ))

	api.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(constants.AuthSignKey)},
	},
	))

	// Initialize private repositories
	ingredientRepo := repositories.NewIngredientRepo(db)
	mealRepo := repositories.NewMealRepository(db)
	measurementRepo := repositories.NewMeasurementRepository(db)

	// Initialize private services
	//userServ - TODO: Implement
	ingredientServ := services.NewIngredientService(ingredientRepo)
	mealServ := services.NewMealService(mealRepo)
	measurementServ := services.NewMeasurementService(measurementRepo)

	// Private handlers

	//	User
	//user := api.Group("/user")
	//user.Get("/", placeholderHandler)
	//user.Post("/", placeholderHandler)
	//user.Get("/:id", placeholderHandler)
	//user.Put("/:id", placeholderHandler)

	// Ingredient
	ingredient := api.Group("/ingredient")
	ingredient.Get("/", IngredientList(ingredientServ))
	ingredient.Post("/", IngredientCreate(ingredientServ, authServ))
	ingredient.Get("/:id", IngredientRetrieve(ingredientServ))
	ingredient.Put("/:id", IngredientUpdate(ingredientServ))

	// Meal
	meal := api.Group("/meal")
	meal.Get("/", MealListByUser(mealServ))
	meal.Get("/all", MealList(mealServ))
	meal.Post("/", MealCreate(mealServ, authServ))
	meal.Get("/:id", MealRetrieve(mealServ))
	meal.Put("/:id", MealUpdate(mealServ))
	meal.Get("/:id/ingredient", MealListIngredients(mealServ))
	meal.Post("/ingredient", MealAddIngredient(mealServ))
	meal.Put("/ingredient/:id", MealIngredientUpdate(mealServ))
	meal.Delete("/ingredient/:id", MealIngredientDelete(mealServ))

	//	Measurement
	measurement := api.Group("/measurement")
	measurement.Get("/", MeasurementListByUserId(measurementServ))
	measurement.Post("/", MeasurementCreate(measurementServ))
	measurement.Post("/:id", MeasurementRetrieve(measurementServ))
	measurement.Put("/:id", MeasurementUpdate(measurementServ))
	measurement.Delete("/:id", MeasurementDelete(measurementServ))
}
