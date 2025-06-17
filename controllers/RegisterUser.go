package controllers

import (
	"encoding/json"
	"time"

	"github.com/NoorUllah43/Chatapp-Backend-in-go.git/db"
	"github.com/NoorUllah43/Chatapp-Backend-in-go.git/middlewares"
	"github.com/NoorUllah43/Chatapp-Backend-in-go.git/models"
	"github.com/gofiber/fiber/v3"
)

func RegisterUser(ctx fiber.Ctx) error {
	var user models.User

	err := json.Unmarshal(ctx.Body(), &user)
	if err != nil {
		return err
	}

	id, err := db.AddUser(user)
	if err != nil {
		return err
	}

	tokenString, err := middlewares.GenerateToken(id)

	cookie := &fiber.Cookie{
		Name:     "chatappToken",
		Value:    tokenString,
		Expires:  time.Now().Add(24 * time.Hour),
	}
	ctx.Cookie(cookie)

	return ctx.JSON("user register successfully ")

}
