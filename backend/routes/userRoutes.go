package routes

import (
	"encoding/json"
	"homefill/backend/auth"
	"homefill/backend/db"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func GetUserInfo(c *fiber.Ctx) error {

	jwtToken := c.Get("Authorization")[7:]
	userId, err := auth.VerifyJwt(jwtToken)
	if err != nil {
		return c.SendStatus(http.StatusUnauthorized)
	}

	user, err := db.GetUserFromId(userId)
	if err != nil {
		return c.SendStatus(http.StatusNotFound)
	}

	data, err := json.Marshal(user)
	if err != nil {
		return c.SendStatus(http.StatusInternalServerError)
	}

	return c.Status(http.StatusOK).Send(data)
}
