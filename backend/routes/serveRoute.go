package routes

import (
	"github.com/gofiber/fiber/v2"
)

// StartServer : start server on port and listen to routes
func StartServer() {
	app := fiber.New()

	app.Get("/", HomeRoute)
	app.Get("/login", LoginRoute)
	app.Get("/callback", AuthCallBack)

	app.Listen(":8080")
}
