package controllers

import (
	"encoding/csv"
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/lorezi/go-admin/database"
	"github.com/lorezi/go-admin/models"
)

func GetOrders(c *fiber.Ctx) error {

	p, _ := strconv.Atoi(c.Query("page", "1"))
	l, _ := strconv.Atoi(c.Query("limit", "5"))

	return c.JSON(models.Paginate(database.DB, &models.Order{}, p, l))
}

func CreateOrder(c *fiber.Ctx) error {
	p := &models.Order{}

	if err := c.BodyParser(&p); err != nil {
		return err
	}

	database.DB.Create(p)

	return c.JSON(p)
}

func GetOrder(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		c.Status(404)
		return c.JSON(fiber.Map{
			"message": "Order not found",
		})
	}
	p := models.Order{
		Id: uint(id),
	}

	database.DB.Find(&p)

	return c.JSON(p)
}

func UpdateOrder(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		c.Status(404)
		return c.JSON(fiber.Map{
			"message": "Order not found",
		})
	}

	p := models.Order{
		Id: uint(id),
	}

	if err := c.BodyParser(&p); err != nil {
		return err
	}

	database.DB.Model(&p).Updates(p)

	return c.JSON(p)

}

func DeleteOrder(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		c.Status(404)
		return c.JSON(fiber.Map{
			"message": "Order not found",
		})
	}

	p := models.Order{
		Id: uint(id),
	}

	database.DB.Delete(&p)

	return nil

}

func Export(c *fiber.Ctx) error {
	filePath := "./csv/orders.csv"
	if err := CreateFile(filePath); err != nil {
		return err
	}

	return c.Download(filePath)
}

func CreateFile(filePath string) error {
	f, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer f.Close()

	w := csv.NewWriter(f)
	defer w.Flush()

	so := []models.Order{}
	database.DB.Preload("OrderItems").Find(&so)

	w.Write([]string{
		"ID", "Name", "Email", "Product Title", "Price", "Quantity",
	})

	for _, v := range so {
		d := []string{
			strconv.Itoa(int(v.Id)),
			v.FirstName + " " + v.LastName,
			v.Email,
			"",
			"",
			"",
		}

		if err := w.Write(d); err != nil {
			return err
		}

		for _, oi := range v.OrderItems {
			// float to string conversion
			p := strconv.FormatFloat(oi.Price, 'f', 2, 64)

			d := []string{
				"",
				"",
				"",
				oi.ProductTitle,
				p,
				strconv.Itoa(int(oi.Quantity)),
			}

			if err := w.Write(d); err != nil {
				return err
			}

		}

	}

	return nil
}

type Sales struct {
	Date string `json:"date"`
	Sum  string `json:"sum"`
}

func Chart(c *fiber.Ctx) error {
	ss := []Sales{}
	database.DB.Raw(
		`SELECT DATE_FORMAT(o.created_at, '%Y-%m-%d') as date, SUM(oi.price * oi.quantity) as sum
		FROM orders o
		JOIN order_items oi on o.id = oi.order_id
		GROUP BY date
		`).Scan(&ss)

	return c.JSON(ss)

}
