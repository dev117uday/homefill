package routes

import (
	"github.com/gofiber/fiber/v2"
)

// StartServer : start server on port and listen to routes
func StartServer() {
	app := fiber.New()

	// Universally Accessible Routes
	app.Get("/", HomeRoute)
	app.Get("/login", LoginRoute)
	app.Get("/callback", AuthCallBack)

	// User Routes : PROTECTED
	app.Get("/user", GetUserInfo)

	app.Listen(":8080")
}
