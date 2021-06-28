package routes

import (
	"go-auth-google/backend/errset"
	"go-auth-google/backend/service"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func GetUserInfo(c *fiber.Ctx) error {

	id, err := service.JwtMiddleWare(c)
	if err != nil {
		return errset.ReturnError(err, c)
	}

	data, err := service.GetUserById(id)
	if err != nil {
		return errset.ReturnError(err, c)
	}

	return c.Status(http.StatusOK).Send(data)
}
