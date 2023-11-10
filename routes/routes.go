package routes

import (
	"sample/handlers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

func SetupRoutes(app *fiber.App) {
	/////////
	app.Get("/rico*", swagger.HandlerDefault)
	////////
	app.Get("/read", handlers.TransCredit)
	app.Get("/read", handlers.TransCredit)
}
