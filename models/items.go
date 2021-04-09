package models

type Item struct {
	Id           uint   `json:"id"`
	OrderId      uint   `json:"order_id"`
	ProductTitle string `json:"product_title"`
	Price        string `json:"price"`
	Quantity     uint   `json:"quantity"`
}
