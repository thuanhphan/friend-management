package model

type User struct {
	Email string `boil:"email" json:"email"`
	Name  string `boil:"name" json:"name"`
}
