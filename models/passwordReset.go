package models

type PasswordReset struct {
	Id    uint
	Email string
	Token string
}
