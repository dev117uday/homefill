package routes

import (
	"encoding/json"
	"homefill/backend/db"
	"homefill/backend/service"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func GetUserInfo(c *fiber.Ctx) error {

	id, err := service.JwtMiddleWare(c)
	if err != nil {
		return err
	}

	user, err := db.GetUserFromId(id)
	if err != nil {
		return err
	}

	data, err := json.Marshal(user)
	if err != nil {
		return c.SendStatus(http.StatusInternalServerError)
	}

	return c.Status(http.StatusOK).Send(data)
}
