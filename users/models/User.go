package models

type User struct {
	Id       uint64
	Name     string
	Email    string
	Password string
}

type CreateUser struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `password:"password"`
}
