package db

import (
	"github.com/NoorUllah43/Chatapp-Backend-in-go.git/middlewares"
	"github.com/NoorUllah43/Chatapp-Backend-in-go.git/models"
)

func AddUser(user models.SignupCredentials) (int, error) {
	query :=
		`INSERT INTO 
	users 
		(name, email, password) 
	VALUES 
		( $1, $2, $3)
	RETURNING id`

	var id int

	hashPassword, err := middlewares.HashPassword(user.Password)
	if err != nil {
		return 0, err
	}

	err = DB.QueryRow(query, user.Name, user.Email, hashPassword).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}
