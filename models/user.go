package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	Id        uint   `json:"id,omitempty"`
	FirstName string `json:"first_name" validate:"required,min=2,max=25"`
	LastName  string `json:"last_name" validate:"required,min=2,max=25"`
	Email     string `json:"email" gorm:"unique" validate:"required,email"`
	// Password        []byte `json:"password"`
	// PasswordConfirm []byte `json:"password_confirm" gorm:"-"`

	Password        string `json:"password" validate:"required"`
	PasswordConfirm string `json:"password_confirm" gorm:"-"`
	RoleId          uint   `json:"role_id" validate:"required"`
	Role            Role   `json:"role" gorm:"foreignKey:RoleId"`
}

type UserResponse struct {
	Id        uint   `json:"id,omitempty"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	RoleId    uint   `json:"role_id"`
	RoleName  string `json:"role_name"`
}

type Login struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func (u *User) SetPassword(p string) {
	hp, _ := bcrypt.GenerateFromPassword([]byte(p), 14)
	u.Password = string(hp)
}

func (u *User) ComparePassword(p string) error {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(p))
	return err
}

func (u *User) Count(db *gorm.DB) int64 {
	var total int64
	db.Model(&User{}).Count(&total)
	return total
}

func (u *User) Paginate(db *gorm.DB, limit int, offset int) interface{} {
	su := []User{}
	db.Preload("Role").Offset(offset).Limit(limit).Find(&su)
	return su
}
