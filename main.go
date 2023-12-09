package main

import (
	"food-backend/src/database"
	"food-backend/src/router"
	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
)

func main() {
	database.ConnectToDB(database.DbConfig)

	app := fiber.New()

	router.SetupRoutes(app)

	err := app.Listen(":3000")
	if err != nil {
		panic(err)
	}
}
