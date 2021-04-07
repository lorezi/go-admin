package controllers

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/lorezi/go-admin/database"
	"github.com/lorezi/go-admin/models"
	"github.com/lorezi/go-admin/util"
	"golang.org/x/crypto/bcrypt"
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

	// TODO create a helper function to generate password
	p, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 12)

	u := models.User{
		FirstName: data["first_name"],
		LastName:  data["last_name"],
		Email:     data["email"],
		Password:  p,
	}

	tx := database.DB.Create(&u)

	fmt.Print(tx)

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

	if err := bcrypt.CompareHashAndPassword(u.Password, []byte(data["password"])); err != nil {
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

type Claims struct {
	jwt.StandardClaims
}

func AuthUser(c *fiber.Ctx) error {
	cookie := c.Cookies("token")

	token, err := util.VerifyJwt(cookie)

	if err != nil || !token.Valid {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unauthenticated user",
		})
	}

	claims := token.Claims.(*Claims)

	u := models.User{}

	database.DB.Where("id = ?", claims.Issuer).First(&u)

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
