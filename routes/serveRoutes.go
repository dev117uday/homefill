package routes

import (
	"fmt"
	config "go-auth-google/backend/config"
	"go-auth-google/backend/logs"

	"github.com/gofiber/fiber/v2"
)

// StartServer : start server on port and listen to routes
func StartServer() {
	app := fiber.New()

	// Universally Accessible Routes
	app.Get("/", HomeRoute)
	app.Get("/login", LoginRoute)
	app.Get("/callback", AuthCallBack)

	// -------------- PROTECTED --------------

	// User Routes
	app.Get("/user", GetUserInfo)

	// ---------------------------------------

	// Starting Server
	err := app.Listen(config.BACKEND_URL)
	if err != nil {
		fmt.Println(err)
		logs.LogIt(logs.LogFatal, "StartServer", "unable to start server", err)
	}
}
