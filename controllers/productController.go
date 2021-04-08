package controllers

import (
	"math"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/lorezi/go-admin/database"
	"github.com/lorezi/go-admin/models"
)

func GetProducts(c *fiber.Ctx) error {
	var total int64
	p, _ := strconv.Atoi(c.Query("page", "1"))
	l, _ := strconv.Atoi(c.Query("limit", "5"))

	o := (p - 1) * l
	sp := []models.Product{}

	database.DB.Offset(o).Limit(l).Find(&sp)

	database.DB.Model(&models.Product{}).Count(&total)

	return c.JSON(fiber.Map{
		"meta": fiber.Map{
			"total":     total,
			"page":      p,
			"last_page": math.Ceil(float64(total) / float64(l)),
		},
		"data": sp,
	})
}

func CreateProduct(c *fiber.Ctx) error {
	p := &models.Product{}

	if err := c.BodyParser(&p); err != nil {
		return err
	}

	database.DB.Create(p)

	return c.JSON(p)
}

func GetProduct(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		c.Status(404)
		return c.JSON(fiber.Map{
			"message": "Product not found",
		})
	}
	p := models.Product{
		Id: uint(id),
	}

	database.DB.Find(&p)

	return c.JSON(p)
}

func UpdateProduct(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		c.Status(404)
		return c.JSON(fiber.Map{
			"message": "Product not found",
		})
	}

	p := models.Product{
		Id: uint(id),
	}

	if err := c.BodyParser(&p); err != nil {
		return err
	}

	database.DB.Model(&p).Updates(p)

	return c.JSON(p)

}

func DeleteProduct(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		c.Status(404)
		return c.JSON(fiber.Map{
			"message": "Product not found",
		})
	}

	p := models.Product{
		Id: uint(id),
	}

	database.DB.Delete(&p)

	return nil

}
