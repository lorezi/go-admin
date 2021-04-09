package models

import "gorm.io/gorm"

type Order struct {
	Id         uint        `json:"id"`
	FirstName  string      `json:"-"`
	LastName   string      `json:"-"`
	Name       string      `json:"name" gorm:"-"`
	Total      float64     `json:"total" gorm:"-"`
	Email      string      `json:"email"`
	UpdatedAt  string      `json:"updated_at"`
	CreatedAt  string      `json:"created_at"`
	OrderItems []OrderItem `json:"order_items" gorm:"foreignKey:OrderId"`
}

type OrderItem struct {
	Id           uint    `json:"id"`
	OrderId      uint    `json:"order_id"`
	ProductTitle string  `json:"product_title"`
	Price        float64 `json:"price"`
	Quantity     uint    `json:"quantity"`
}

func (p *Order) Count(db *gorm.DB) int64 {
	var total int64
	db.Model(&Order{}).Count(&total)
	return total
}

func (p *Order) Paginate(db *gorm.DB, limit int, offset int) interface{} {
	so := []Order{}
	db.Preload("OrderItems").Offset(offset).Limit(limit).Find(&so)

	for i := range so {
		so[i].Name = so[i].FirstName + " " + so[i].LastName
		total := 0.0
		for _, v := range so[i].OrderItems {
			total += v.Price * float64(v.Quantity)
			so[i].Total = total
		}
	}
	return so
}
