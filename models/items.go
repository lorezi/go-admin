package models

import "gorm.io/gorm"

type Item struct {
	Id           uint   `json:"id"`
	ItemId       uint   `json:"order_id"`
	ProductTitle string `json:"product_title"`
	Price        string `json:"price"`
	Quantity     uint   `json:"quantity"`
}

func (p *Item) Count(db *gorm.DB) int64 {
	var total int64
	db.Model(&Item{}).Count(&total)
	return total
}

func (p *Item) Paginate(db *gorm.DB, limit int, offset int) interface{} {
	sp := []Item{}
	db.Offset(offset).Limit(limit).Find(&sp)
	return sp
}
