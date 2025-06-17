package db


func FindUser(email string) (string, int, error) {
	var password string
	var id int

	user := `SELECT password, id FROM users WHERE email=$1`

	err := DB.QueryRow(user, email).Scan(&password, &id)
	if err != nil {
		return "", 0, err
	}

	return password, id, nil
}