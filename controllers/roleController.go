package controllers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/lorezi/go-admin/database"
	"github.com/lorezi/go-admin/models"
	"gorm.io/gorm/clause"
)

type RoleCreateDTO struct {
	Name        string   `json:"name"`
	Permissions []string `json:"permissions"`
}

func Roles(c *fiber.Ctx) error {
	sr := []models.Role{}

	database.DB.Preload("Permissions").Find(&sr)

	return c.JSON(sr)
}

func CreateRole(c *fiber.Ctx) error {
	rDTO := RoleCreateDTO{}

	if err := c.BodyParser(&rDTO); err != nil {
		return err
	}

	p := make([]models.Permission, len(rDTO.Permissions))

	for i, v := range rDTO.Permissions {
		id, _ := strconv.Atoi(v)
		p[i] = models.Permission{
			Id: uint(id),
		}
	}

	r := models.Role{
		Name:        rDTO.Name,
		Permissions: p,
	}

	database.DB.Create(&r)

	return c.JSON(r)
}

func GetRole(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		c.Status(404)
		return c.JSON(fiber.Map{
			"message": "role not found",
		})
	}
	r := models.Role{
		Id: uint(id),
	}

	database.DB.Preload("Permissions").Find(&r)

	return c.JSON(r)
}

func UpdateRole(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		c.Status(404)
		return c.JSON(fiber.Map{
			"message": "role not found",
		})
	}

	rDTO := RoleCreateDTO{}

	if err := c.BodyParser(&rDTO); err != nil {
		return err

	}

	p := make([]models.Permission, len(rDTO.Permissions))

	for i, v := range rDTO.Permissions {
		id, _ := strconv.Atoi(v)
		p[i] = models.Permission{
			Id: uint(id),
		}
	}

	// Delete previous permissions
	database.DB.Table("role_permissions").Where("role_id", id).Delete(nil)

	r := models.Role{
		Id:          uint(id),
		Name:        rDTO.Name,
		Permissions: p,
	}

	database.DB.Model(&r).Updates(r)

	return c.JSON(r)

}

func DeleteRole(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		c.Status(404)
		return c.JSON(fiber.Map{
			"message": "role not found",
		})
	}

	r := models.Role{
		Id: uint(id),
	}

	database.DB.Select(clause.Associations).Delete(&r)

	return nil

}
