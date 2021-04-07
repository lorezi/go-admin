package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lorezi/go-admin/database"
	"github.com/lorezi/go-admin/models"
)

func GetPermissions(c *fiber.Ctx) error {
	p := models.Permission{}

	database.DB.Find(&p)

	return c.JSON(p)
}
