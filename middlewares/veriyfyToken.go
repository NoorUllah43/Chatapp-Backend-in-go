package middlewares

import (
	"fmt"
	"os"

	"github.com/NoorUllah43/Chatapp-Backend-in-go.git/models"
	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v5"
)

func VerifyToken(ctx fiber.Ctx) error {
	tokenString := ctx.Cookies("chatappToken")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		return ctx.Status(404).JSON(models.ErrorResponse{Success: false, Message: "invalid token"})
	}
	if !token.Valid {
		return ctx.Status(404).JSON(models.ErrorResponse{Success: false, Message: "invalid token"})
	}

	claims := token.Claims.(jwt.MapClaims)

	floatID, ok := claims["userID"].(float64)
	if !ok {
		return ctx.Status(404).JSON(models.ErrorResponse{Success: false, Message: "user id not found"})
	}

	userID := int(floatID)

	if userID <= 0 {
		return ctx.Status(404).JSON(models.ErrorResponse{Success: false, Message: "invalid user id"})
	}
	fmt.Println("userID:", userID)
	ctx.Locals("userID", userID)

	return ctx.Next()
}
