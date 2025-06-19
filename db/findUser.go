package db

import (
	"fmt"

	"github.com/NoorUllah43/Chatapp-Backend-in-go.git/models"
)

func FindUser(email string) (string, string, int, error) {
	var password string
	var id int
	var name string

	user := `SELECT name, password, id FROM users WHERE email=$1`

	err := DB.QueryRow(user, email).Scan(&name, &password, &id)
	if err != nil {
		return "", "", 0, err
	}

	return name, password, id, nil
}

func IsUserExist(email string) bool {
	var exist bool

	query := `SELECT EXISTS(SELECT 1 FROM users WHERE email = $1)`
	err := DB.QueryRow(query, email).Scan(&exist)
	if err != nil {
		fmt.Println("error in query", err)
	}

	return exist

}

func GetAllUsers() ([]models.UserData, error) {

	query := `SELECT name, email, id FROM users`

	rows, err := DB.Query(query)
	if err != nil {
		return nil, err
	}

	var AllUsers []models.UserData

	for rows.Next() {
		var user models.UserData

		err := rows.Scan(&user.Name, &user.Email, &user.ID)
		if err != nil {
			return nil, err
		}
		AllUsers = append(AllUsers, user)
	}

	return AllUsers, nil

}
