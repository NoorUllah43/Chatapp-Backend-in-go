package db

import (

	"github.com/NoorUllah43/Chatapp-Backend-in-go.git/models"
)

func AddUser(user models.User) (int, error) {
	query :=
		`INSERT INTO 
	users 
		(name, email, password) 
	VALUES 
		( $1, $2, $3)
	RETURNING id`

	var id int

	err := DB.QueryRow(query, user.Name, user.Email, user.Password).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}
