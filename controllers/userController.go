package controllers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/lorezi/go-admin/database"
	"github.com/lorezi/go-admin/middlewares"
	"github.com/lorezi/go-admin/models"
)

// TODO add authorization to other routes

func Users(c *fiber.Ctx) error {

	if err := middlewares.IsAuthorized(c, "users"); err != nil {
		return err
	}

	p, _ := strconv.Atoi(c.Query("page", "1"))
	l, _ := strconv.Atoi(c.Query("limit", "5"))

	return c.JSON(models.Paginate(database.DB, &models.User{}, p, l))
}

func CreateUser(c *fiber.Ctx) error {
	if err := middlewares.IsAuthorized(c, "users"); err != nil {
		return err
	}

	u := &models.User{}

	if err := c.BodyParser(&u); err != nil {
		return err
	}

	u.SetPassword("12345")
	//
	// u.RoleId = 1

	database.DB.Create(u)

	return c.JSON(u)
}

func GetUser(c *fiber.Ctx) error {
	if err := middlewares.IsAuthorized(c, "users"); err != nil {
		return err
	}

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
	if err := middlewares.IsAuthorized(c, "users"); err != nil {
		return err
	}

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
	if err := middlewares.IsAuthorized(c, "users"); err != nil {
		return err
	}

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
