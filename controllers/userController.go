package controllers

import (
	"math"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/lorezi/go-admin/database"
	"github.com/lorezi/go-admin/models"
)

func Users(c *fiber.Ctx) error {
	var total int64
	p, _ := strconv.Atoi(c.Query("page", "1"))
	l, _ := strconv.Atoi(c.Query("limit", "5"))

	o := (p - 1) * l
	su := []models.User{}

	database.DB.Preload("Role").Offset(o).Limit(l).Find(&su)

	database.DB.Model(&models.User{}).Count(&total)

	return c.JSON(fiber.Map{
		"meta": fiber.Map{
			"total":     total,
			"page":      p,
			"last_page": math.Ceil(float64(total) / float64(l)),
		},
		"data": su,
	})
}

func CreateUser(c *fiber.Ctx) error {
	u := models.User{}

	if err := c.BodyParser(&u); err != nil {
		return err
	}

	u.SetPassword("12345")
	//
	// u.RoleId = 1

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
