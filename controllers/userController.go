package controllers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/lorezi/go-admin/database"
	"github.com/lorezi/go-admin/models"
)

func Users(c *fiber.Ctx) error {
	su := []models.User{}

	database.DB.Preload("Role").Find(&su)

	return c.JSON(su)
}

func CreateUser(c *fiber.Ctx) error {
	u := models.User{}

	if err := c.BodyParser(&u); err != nil {
		return err
	}

	u.SetPassword("12345")

	database.DB.Create(&u)

	return c.JSON(u)
}

func GetUser(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		c.Status(404)
		return c.JSON(fiber.Map{
			"message": "user not found",
		})
	}
	u := models.User{
		Id: uint(id),
	}

	database.DB.Preload("Role").Find(&u)

	return c.JSON(u)
}

func UpdateUser(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		c.Status(404)
		return c.JSON(fiber.Map{
			"message": "user not found",
		})
	}

	u := models.User{
		Id: uint(id),
	}

	if err := c.BodyParser(&u); err != nil {
		return err
	}

	database.DB.Model(&u).Updates(u)

	return c.JSON(u)

}

func DeleteUser(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		c.Status(404)
		return c.JSON(fiber.Map{
			"message": "user not found",
		})
	}

	u := models.User{
		Id: uint(id),
	}

	database.DB.Delete(&u)

	return nil

}
