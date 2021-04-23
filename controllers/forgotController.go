package controllers

import (
	"fmt"
	"math/rand"
	"net/smtp"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/lorezi/go-admin/database"
	"github.com/lorezi/go-admin/models"
	"github.com/subosito/gotenv"
	"golang.org/x/crypto/bcrypt"
)

func Forgot(c *fiber.Ctx) error {
	gotenv.Load()
	u := new(models.PasswordReset)

	if err := c.BodyParser(&u); err != nil {
		return err
	}

	token := tokenGenerator(12)

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

	// token expiration time is 3hr
	pr.ExpirationTime = time.Now().Add(time.Hour * time.Duration(3))
	pr.CreatedAt = time.Now()

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

func ResetPassword(c *fiber.Ctx) error {

	rp := &models.PasswordReset{}

	if err := database.DB.Where("token = ?", c.Params("token")).Last(rp); err.Error != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "invalid token",
		})
	}

	if rp.Id == 0 {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "invalid token",
		})
	}

	now, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

	if now.After(rp.ExpirationTime) {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "token has expired",
		})
	}

	r := new(models.Reset)

	if r.Password != r.PasswordConfirm {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "password does not match",
		})
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(r.Password), 14)
	database.DB.Model(&models.User{}).Where("email = ?", rp.Email).Update("password", password)

	return c.JSON(fiber.Map{
		"message": "success",
	})

}

func tokenGenerator(n int) string {
	b := make([]byte, n)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}
