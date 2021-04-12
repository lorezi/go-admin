package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	Id        uint   `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email" gorm:"unique"`
	Password  []byte `json:"-"`
	RoleId    uint   `json:"role_id"`
	Role      Role   `json:"role" gorm:"foreignKey:RoleId"`
}

func (u *User) SetPassword(p string) {
	// default password for created users
	hp, _ := bcrypt.GenerateFromPassword([]byte(p), 14)
	u.Password = hp
}

func (u *User) ComparePassword(p string) error {
	err := bcrypt.CompareHashAndPassword(u.Password, []byte(p))
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
