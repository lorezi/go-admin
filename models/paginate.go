package models

import (
	"math"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Paginate(db *gorm.DB, p int, l int) fiber.Map {
	var total int64

	o := (p - 1) * l
	sp := []Product{}

	db.Offset(o).Limit(l).Find(&sp)

	db.Model(&Product{}).Count(&total)

	return fiber.Map{
		"meta": fiber.Map{
			"total":     total,
			"page":      p,
			"last_page": math.Ceil(float64(total) / float64(l)),
		},
		"data": sp,
	}
}
