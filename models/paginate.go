package models

import (
	"math"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Paginate(db *gorm.DB, md CountPaginate, p int, l int) fiber.Map {

	o := (p - 1) * l

	data := md.Paginate(db, l, o)
	total := md.Count(db)
	return fiber.Map{
		"meta": fiber.Map{
			"total":     total,
			"page":      p,
			"last_page": math.Ceil(float64(total) / float64(l)),
		},
		"data": data,
	}
}
