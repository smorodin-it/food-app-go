package handlers

import (
	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
)

func SetupTestApp() (app *fiber.App) {
	app = fiber.New()
	SetupApp(app)

	return app
}
