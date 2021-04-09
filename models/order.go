package models

type Order struct {
	Id         uint   `json:"id"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Email      string `json:"email"`
	UpdatedAt  string `json:"updated_at"`
	CreatedAt  string `json:"created_at"`
	OrderItems []Item `json:"order_items" gorm:"foreignKey:OrderId"`
}
