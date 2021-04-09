package models

import "gorm.io/gorm"

type Product struct {
	Id          uint    `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Image       string  `json:"image"`
	Price       float64 `json:"price"`
}

func (p *Product) Count(db *gorm.DB) int64 {
	var total int64
	db.Model(&Product{}).Count(&total)
	return total
}

func (p *Product) Paginate(db *gorm.DB, limit int, offset int) interface{} {
	sp := []Product{}
	db.Offset(offset).Limit(limit).Find(&sp)
	return sp
}
