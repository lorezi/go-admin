package controllers

import (
	"os"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/lorezi/go-admin/database"
	"github.com/lorezi/go-admin/models"
	"github.com/lorezi/go-admin/util"
)

var SECRET_KEY string = os.Getenv("SECRET_KEY")

func Register(c *fiber.Ctx) error {

	// TODO Add validation
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	if data["password"] != data["password_confirm"] {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "passwords do not match",
		})
	}

	u := models.User{
		FirstName: data["first_name"],
		LastName:  data["last_name"],
		Email:     data["email"],
		RoleId:    1,
	}
	u.SetPassword(data["password"])

	database.DB.Create(&u)

	return c.JSON(u)
}

func Login(c *fiber.Ctx) error {

	data := make(map[string]string)

	u := models.User{}

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	database.DB.Where("email = ?", data["email"]).First(&u)

	if u.Id == 0 {
		c.Status(404)
		return c.JSON(fiber.Map{
			"message": "invalid login credentials 😰",
		})
	}

	if err := u.ComparePassword(data["password"]); err != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "invalid login credentials 😰",
		})
	}

	token, err := util.GenerateJwt(strconv.Itoa(int(u.Id)))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	cookie := fiber.Cookie{
		Name:     "token",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24), //1 day ,
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "success",
	})

}

func AuthUser(c *fiber.Ctx) error {

	cookie := c.Cookies("token")

	userId, _ := util.VerifyJwt(cookie)

	u := models.User{}

	database.DB.Where("id = ?", userId).First(&u)

	return c.JSON(u)
}

func Logout(c *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name:     "token",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour), //1 day ,
		HTTPOnly: true,
	}
	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "success",
	})

}

// User bioprofile
func UpdateInfo(c *fiber.Ctx) error {
	data := make(map[string]string)

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	cookie := c.Cookies("token")

	Id, _ := util.VerifyJwt(cookie)

	userId, _ := strconv.Atoi(Id)
	u := models.User{
		Id:        uint(userId),
		FirstName: data["first_name"],
		LastName:  data["last_name"],
		Email:     data["email"],
	}

	database.DB.Model(&u).Updates(u)

	return c.JSON(u)

}

func UpdatePassword(c *fiber.Ctx) error {
	data := make(map[string]string)

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	if data["password"] != data["password_confirm"] {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "passwords do not match",
		})
	}

	cookie := c.Cookies("token")

	Id, _ := util.VerifyJwt(cookie)
	userId, _ := strconv.Atoi(Id)

	u := &models.User{
		Id: uint(userId),
	}

	u.SetPassword(data["password"])

	database.DB.Model(&u).Updates(u)

	// successful update remove cookies
	rmCookie := fiber.Cookie{
		Name:     "token",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour), //1 day ,
		HTTPOnly: true,
	}
	c.Cookie(&rmCookie)

	return c.JSON(u)

}
