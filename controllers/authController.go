package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lorezi/go-admin/models"
)

func Register(c *fiber.Ctx) error {
	u := models.User{
		FirstName: "John",
	}

	u.LastName = "Doe"

	return c.JSON(u)
}
