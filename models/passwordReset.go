package models

type PasswordReset struct {
	Id    uint   `json:"-"`
	Email string `json:"email" validate:"required,email"`
	Token string `json:"-"`
}
