package middlewares

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(userID int) (string, error) {

	jwtSecret := []byte(os.Getenv("JWT_SECRET"))

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID": userID,
		"exp":       time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenstring, err := token.SignedString(jwtSecret)

	return tokenstring, err

}
