package controllers

import "github.com/gofiber/fiber/v2"

func Upload(c *fiber.Ctx) error {
	form, err := c.MultipartForm()

	if err != nil {
		return err
	}

	sf := form.File["image"]
	fn := ""

	for _, f := range sf {
		fn = f.Filename

		if err := c.SaveFile(f, "./uploads/"+fn); err != nil {
			return err
		}
	}

	return c.JSON(fiber.Map{
		"url": "http://localhost:8080/api/uploads/" + fn,
	})

}
