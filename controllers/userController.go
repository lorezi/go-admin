package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lorezi/go-admin/database"
	"github.com/lorezi/go-admin/models"
	"golang.org/x/crypto/bcrypt"
)

func Users(c *fiber.Ctx) error {
	su := []models.User{}

	database.DB.Find(&su)

	return c.JSON(su)
}

func CreateUser(c *fiber.Ctx) error {
	u := models.User{}

	if err := c.BodyParser(&u); err != nil {
		return err
	}

	// default password for created users
	p, _ := bcrypt.GenerateFromPassword([]byte("123456"), 14)
	u.Password = p

	database.DB.Create(&u)

	return c.JSON(u)
}
