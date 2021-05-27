package routes

import (
	"fmt"
	"homefill/backend/auth"
	config "homefill/backend/configs"
	"homefill/backend/db"
	"homefill/backend/functions"
	"homefill/backend/model"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

// StartServer : start server on port and listen to routes
func StartServer() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		db.GetUserFromId(`112537321089373511651`)
		return c.SendString("hello world")
	})

	app.Get("/login", func(c *fiber.Ctx) error {
		url := config.GOOGLEAuthConfig.AuthCodeURL(config.State)
		return c.Redirect(url, http.StatusTemporaryRedirect)
	})

	app.Get("/callback", func(c *fiber.Ctx) error {
		content, err := auth.GetUserInfo(c.FormValue("state"), c.FormValue("code"))
		if err != nil {
			fmt.Println(err.Error())
			return c.Redirect(config.FRONTEND_URL, http.StatusUnauthorized)
		}

		var user model.User
		functions.ConvertToStruct(content, &user)

		// TODO : if user exists, generate jwt token and return
		_, ans := db.GetUserFromId(user.ID)
		if ans {
			token, err := auth.GenerateJwtToken(user.ID)
			fmt.Println(token)
			if err != nil {
				fmt.Print(err)
				return c.SendStatus(http.StatusUnauthorized)
			} else {
				return c.SendString(token)
			}
		} else {
			// TODO : if user doesn't exists, create user, generate jwt
			answer := db.InsertUser(&user)
			if answer {
				token, err := auth.GenerateJwtToken(user.ID)
				fmt.Println(token)
				if err != nil {
					fmt.Print(err)
					return c.SendStatus(http.StatusUnauthorized)
				} else {
					return c.SendString(token)
				}
			} else {
				return c.SendStatus(http.StatusUnauthorized)
			}
		}

	})

	app.Listen(":8080")
}
