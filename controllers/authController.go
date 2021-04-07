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
	"golang.org/x/crypto/bcrypt"
)

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

	SECRET_KEY := os.Getenv("SECRET_KEY")
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

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(u.Id)),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), // 1 day
	})

	token, err := claims.SignedString([]byte(SECRET_KEY))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(token)

}