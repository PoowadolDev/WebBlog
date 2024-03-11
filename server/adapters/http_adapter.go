package adapters

import (
	"Web_Blog/core/domain"
	"Web_Blog/core/service"

	"github.com/gofiber/fiber/v2"
)

type HttpUserhandler struct {
	service service.UserService
}

func NewHttpUserHandler(service service.UserService) *HttpUserhandler {
	return &HttpUserhandler{service: service}
}

func (h *HttpUserhandler) CreateUser(c *fiber.Ctx) error {
	var user domain.User
	err := c.BodyParser(&user)
	if err != nil {
		return err
	}

	user.Role = "user"

	err = h.service.CreateUser(user)
	if err != nil {
		return err
	}
	return c.JSON(user)
}
