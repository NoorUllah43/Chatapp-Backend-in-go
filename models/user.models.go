package models

type SignupCredentials struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginCredentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}


type UserData struct {
	ID  int 
	Name     string `json:"name"`
	Email    string `json:"email"`
}