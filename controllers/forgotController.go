package controllers

import (
	"math/rand"
	"net/smtp"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/lorezi/go-admin/database"
	"github.com/lorezi/go-admin/models"
	"github.com/subosito/gotenv"
)

func Forgot(c *fiber.Ctx) error {
	gotenv.Load()
	u := new(models.PasswordReset)

	if err := c.BodyParser(&u); err != nil {
		return err
	}

	token := randStringByte(12)

	pr := &models.PasswordReset{
		Email: u.Email,
		Token: token,
	}

	// search for the email in the database, if the user exist
	um := &models.User{}

	database.DB.Where("email = ?", u.Email).First(um)

	if um.Id == 0 {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "invalid email address 😰",
		})
	}

	database.DB.Create(pr)

	from := os.Getenv("EMAIL_FROM")

	to := []string{
		u.Email,
	}

	auth := smtp.PlainAuth("", os.Getenv("EMAIL_USERNAME"), os.Getenv("EMAIL_PASSWORD"), os.Getenv("EMAIL_HOST"))

	url := os.Getenv("RESET_URL") + token

	msg := []byte("Click <a href=\"" + url + "\">here</a> to reset your password!")

	err := smtp.SendMail(os.Getenv("EMAIL_HOST")+":"+os.Getenv("EMAIL_PORT"), auth, from, to, msg)
	if err != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "email was not sent 😰",
		})
	}

	return c.JSON(fiber.Map{
		"message": "success",
	})

}

func randStringByte(n int) string {
	lb := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, n)
	for i := range b {
		b[i] = lb[rand.Intn(len(lb))]
	}
	return string(b)
}
