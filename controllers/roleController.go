package controllers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/lorezi/go-admin/database"
	"github.com/lorezi/go-admin/models"
)

func Roles(c *fiber.Ctx) error {
	sr := []models.Role{}

	database.DB.Find(&sr)

	return c.JSON(sr)
}

func CreateRole(c *fiber.Ctx) error {
	r := models.Role{}

	if err := c.BodyParser(&r); err != nil {
		return err
	}

	database.DB.Create(&r)

	return c.JSON(r)
}

func GetRole(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		c.Status(404)
		return c.JSON(fiber.Map{
			"message": "user not found",
		})
	}
	r := models.Role{
		Id: uint(id),
	}

	database.DB.Find(&r)

	return c.JSON(r)
}

func UpdateRole(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		c.Status(404)
		return c.JSON(fiber.Map{
			"message": "user not found",
		})
	}

	r := models.Role{
		Id: uint(id),
	}

	if err := c.BodyParser(&r); err != nil {
		return err
	}

	database.DB.Model(&r).Updates(r)

	return c.JSON(r)

}

func DeleteRole(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		c.Status(404)
		return c.JSON(fiber.Map{
			"message": "user not found",
		})
	}

	r := models.User{
		Id: uint(id),
	}

	database.DB.Delete(&r)

	return nil

}
