package db

import (
	"github.com/NoorUllah43/Chatapp-Backend-in-go.git/middlewares"
	"github.com/NoorUllah43/Chatapp-Backend-in-go.git/models"
)

func AddUser(user models.SignupCredentials) (int, error) {
	createTableQuery := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		name VARCHAR(100) NOT NULL,
		email VARCHAR(100) UNIQUE NOT NULL,
		password VARCHAR(255) NOT NULL
	);`

	insertUser :=
		`
		INSERT INTO 
			users (name, email, password) 
			VALUES ( $1, $2, $3)
			RETURNING id;`

	var id int

	hashPassword, err := middlewares.HashPassword(user.Password)
	if err != nil {
		return 0, err
	}
	_, err = DB.Exec(createTableQuery)
	if err != nil {
		return 0, err
	}

	err = DB.QueryRow(insertUser, user.Name, user.Email, hashPassword).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}
