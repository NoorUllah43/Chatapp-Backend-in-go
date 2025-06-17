package models

type User struct {
	Name     string `json:name`
	Email    string `json:email`
	password string `json:password`
}

type LoginCredentials struct {
	Email    string `json:email`
	password string `json:password`
}
