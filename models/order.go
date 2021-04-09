package models

import "gorm.io/gorm"

type Order struct {
	Id         uint   `json:"id"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Email      string `json:"email"`
	UpdatedAt  string `json:"updated_at"`
	CreatedAt  string `json:"created_at"`
	OrderItems []Item `json:"order_items" gorm:"foreignKey:OrderId"`
}

func (p *Order) Count(db *gorm.DB) int64 {
	var total int64
	db.Model(&Order{}).Count(&total)
	return total
}

func (p *Order) Paginate(db *gorm.DB, limit int, offset int) interface{} {
	sp := []Order{}
	db.Offset(offset).Limit(limit).Find(&sp)
	return sp
}
