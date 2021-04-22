package controllers

import (
	"math/rand"
	"net/smtp"

	"github.com/gofiber/fiber/v2"
	"github.com/lorezi/go-admin/database"
	"github.com/lorezi/go-admin/models"
)

func Forgot(c *fiber.Ctx) error {
	u := new(models.PasswordReset)

	if err := c.BodyParser(&u); err != nil {
		return err
	}

	token := RandStringRunes(12)

	pr := &models.PasswordReset{
		Email: u.Email,
		Token: token,
	}

	database.DB.Create(pr)

	// use env variable here
	from := "admin@example.com"

	to := []string{
		u.Email,
	}

	// env variable
	url := "http://localhost:3000/reset/" + token

	// env variable
	m := []byte("Click <a href=\"" + url + "\">here</a> to reset your password!")

	err := smtp.SendMail("0.0.0.0:1025", nil, from, to, m)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"message": "success",
	})

}

func RandStringRunes(n int) string {
	lr := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		b[i] = lr[rand.Intn(len(lr))]
	}
	return string(b)
}
