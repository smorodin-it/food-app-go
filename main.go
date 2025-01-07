package main

import (
	_ "food-backend/docs"
	"food-backend/src/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	_ "github.com/lib/pq"
)

// @Title food-app backend
// @Version 1.0
// @Description swagger for food-app api
// @Host 127.0.0.1:3000
// @BasePath /api
// @SecurityDefinitions.apikey ApiKeyAuth
// @In header
// @Name Authorization
func main() {
	app := fiber.New()

	app.Get("/swagger/*", swagger.HandlerDefault)

	router.SetupRoutes(app)

	err := app.Listen(":3000")
	if err != nil {
		panic(err)
	}
}
