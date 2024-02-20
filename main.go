package main

import (
	_ "food-backend/docs"
	"food-backend/src/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	_ "github.com/lib/pq"
)

// @title food-app backend
// @version 1.0
// @description swagger for food-app api
// @host 127.0.0.1:3000
// @BasePath /api
func main() {
	app := fiber.New()

	app.Get("/swagger/*", swagger.HandlerDefault)

	router.SetupRoutes(app)

	err := app.Listen(":3000")
	if err != nil {
		panic(err)
	}
}
