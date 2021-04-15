package models

type Role struct {
	Id          uint         `json:"id" gorm:"index:,unique"`
	Name        string       `json:"name"`
	Permissions []Permission `json:"permissions" gorm:"many2many:role_permissions;constraint:OnDelete:CASCADE"`
}
