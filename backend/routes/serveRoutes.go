package routes

import (
	config "homefill/backend/config"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
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
	err := app.Listen(config.PORT)
	if err != nil {
		config.Log.WithFields(logrus.Fields{
			"fn":  "StartServer",
			"err": err.Error(),
		}).Fatal("unable to start server")
	}
}
