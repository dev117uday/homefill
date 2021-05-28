package routes

import (
	"encoding/json"
	"homefill/backend/auth"
	config "homefill/backend/configs"
	"homefill/backend/db"
	"homefill/backend/model"
	"homefill/backend/service"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func HomeRoute(c *fiber.Ctx) error {

	// TODO : proper error handling in getuserfromid func
	user, err := db.GetUserFromId(`118366008323505800389`)

	// TODO : proper error handling when return

	if err != nil {
		return c.SendStatus(http.StatusNotFound)
	}

	data, err := json.Marshal(user)
	if err != nil {
		return c.SendStatus(http.StatusInternalServerError)
	}

	return c.Status(http.StatusOK).Send(data)
}

func LoginRoute(c *fiber.Ctx) error {
	url := config.GOOGLEAuthConfig.AuthCodeURL(config.State)
	return c.Redirect(url, http.StatusTemporaryRedirect)
}

func AuthCallBack(c *fiber.Ctx) error {

	content, err := auth.GetUserInfo(c.FormValue("state"), c.FormValue("code"))
	if err != nil {
		return c.SendStatus(http.StatusInternalServerError)
	}

	var user model.User
	json.Unmarshal(content, &user)
	token, err := service.GenerateJwtTokenService(&user)

	if err != nil {
		return c.SendStatus(http.StatusInternalServerError)
	}

	return c.Status(http.StatusOK).SendString(token)
}
