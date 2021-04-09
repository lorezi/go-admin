package controllers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/lorezi/go-admin/database"
	"github.com/lorezi/go-admin/models"
)

func GetOrders(c *fiber.Ctx) error {

	p, _ := strconv.Atoi(c.Query("page", "1"))
	l, _ := strconv.Atoi(c.Query("limit", "5"))

	return c.JSON(models.Paginate(database.DB, &models.Order{}, p, l))
}

func CreateOrder(c *fiber.Ctx) error {
	p := &models.Order{}

	if err := c.BodyParser(&p); err != nil {
		return err
	}

	database.DB.Create(p)

	return c.JSON(p)
}

func GetOrder(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		c.Status(404)
		return c.JSON(fiber.Map{
			"message": "Order not found",
		})
	}
	p := models.Order{
		Id: uint(id),
	}

	database.DB.Find(&p)

	return c.JSON(p)
}

func UpdateOrder(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		c.Status(404)
		return c.JSON(fiber.Map{
			"message": "Order not found",
		})
	}

	p := models.Order{
		Id: uint(id),
	}

	if err := c.BodyParser(&p); err != nil {
		return err
	}

	database.DB.Model(&p).Updates(p)

	return c.JSON(p)

}

func DeleteOrder(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		c.Status(404)
		return c.JSON(fiber.Map{
			"message": "Order not found",
		})
	}

	p := models.Order{
		Id: uint(id),
	}

	database.DB.Delete(&p)

	return nil

}
