package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lorezi/go-admin/database"
	"github.com/lorezi/go-admin/models"
)

func Users(c *fiber.Ctx) error {
	su := []models.User{}

	database.DB.Find(&su)

	return c.JSON(su)
}
