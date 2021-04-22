package controllers

import (
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/lorezi/go-admin/database"
	"github.com/lorezi/go-admin/models"
	"github.com/lorezi/go-admin/util"
)

var SECRET_KEY string = os.Getenv("SECRET_KEY")

func Register(c *fiber.Ctx) error {

	nu := new(models.User)

	if err := c.BodyParser(&nu); err != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if nu.Password != nu.PasswordConfirm {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "passwords do not match",
		})
	}

	u := &models.User{
		FirstName: nu.FirstName,
		LastName:  nu.LastName,
		Email:     nu.Email,
		RoleId:    6,
	}

	u.SetPassword(nu.Password)

	if err := util.ValidateStruct(*u); err != nil {
		c.Status(400)
		return c.JSON(err)
	}

	if err := database.DB.Create(u).Error; err != nil {
		c.Status(500)
		sm := strings.Split(err.Error(), ":")
		m := strings.TrimSpace(sm[1])

		return c.JSON(fiber.Map{
			"message": m,
		})
	}

	return c.JSON(fiber.Map{
		"message": "user account created",
	})
}

func Login(c *fiber.Ctx) error {

	lu := new(models.Login)

	if err := c.BodyParser(&lu); err != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if err := util.ValidateStruct(*lu); err != nil {
		c.Status(400)
		return c.JSON(err)
	}

	u := &models.User{}

	database.DB.Where("email = ?", lu.Email).First(&u)

	if u.Id == 0 {
		c.Status(404)
		return c.JSON(fiber.Map{
			"message": "invalid login credentials email 😰",
		})
	}

	if err := u.ComparePassword(lu.Password); err != nil {

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

	database.DB.Where("id = ?", userId).Preload("Role").First(&u)

	r := &models.UserResponse{
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Email:     u.Email,
		RoleId:    u.Role.Id,
		RoleName:  u.Role.Name,
	}

	return c.JSON(r)
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
