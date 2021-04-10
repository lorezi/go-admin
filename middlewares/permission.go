package middlewares

import (
	"errors"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/lorezi/go-admin/database"
	"github.com/lorezi/go-admin/models"
	"github.com/lorezi/go-admin/util"
)

func IsAuthorized(c *fiber.Ctx, page string) error {

	cookie := c.Cookies("token")
	id, err := util.VerifyJwt(cookie)
	if err != nil {
		return err
	}

	userId, err := strconv.Atoi(id)
	if err != nil {
		return err
	}

	u := models.User{
		Id: uint(userId),
	}

	database.DB.Preload("Role").Find(&u)

	r := models.Role{
		Id: u.RoleId,
	}

	database.DB.Preload("Permissions").Find(&r)

	if c.Method() == "GET" {
		for _, permission := range r.Permissions {

			if permission.Name == "view_"+page || permission.Name == "edit_"+page {
				return nil
			}
		}
	} else {
		for _, permission := range r.Permissions {
			if permission.Name == "edit_"+page {
				return nil
			}
		}
	}

	c.Status(fiber.StatusUnauthorized)
	return errors.New("unauthorized")

}
