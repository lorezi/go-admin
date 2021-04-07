package models

import "golang.org/x/crypto/bcrypt"

type User struct {
	Id        uint   `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email" gorm:"unique"`
	Password  []byte `json:"-"`
	RoleId    uint   `json:"role_id"`
	Role      Role   `json:"role" gorm:"foreignKey:RoleId"`
}

func (u User) SetPassword(p string) {
	// default password for created users
	hp, _ := bcrypt.GenerateFromPassword([]byte(p), 14)
	u.Password = hp
}

func (u User) ComparePassword(p string) error {
	err := bcrypt.CompareHashAndPassword(u.Password, []byte(p))
	return err
}
