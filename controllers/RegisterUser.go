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
	var user models.SignupCredentials

	err := json.Unmarshal(ctx.Body(), &user)
	if err != nil {
		ctx.JSON(models.ErrorResponse{Success: false, Message: "provide correct credentials"})
	}

	// check for values
	if user.Email == "" || user.Name == "" || user.Password == "" {
		return ctx.JSON(models.ErrorResponse{Success: false, Message: "provide name, email and password"})
	}

	// check is user exist
	exist := db.IsUserExist(user.Email)
	if exist {
		return ctx.JSON(models.ErrorResponse{Success: false, Message: "user already exist"})
	}

	id, err := db.AddUser(user)
	if err != nil {
		return ctx.JSON(models.ErrorResponse{Success: false, Message: "provide correct credentials"})
	}

	tokenString, err := middlewares.GenerateToken(id)
	if err != nil {
		ctx.JSON(models.ErrorResponse{Success: false, Message: "error in generating token"})
	}

	cookie := &fiber.Cookie{
		Name:    "chatappToken",
		Value:   tokenString,
		Expires: time.Now().Add(24 * time.Hour),
	}
	ctx.Cookie(cookie)

	return ctx.Status(201).JSON(models.SuccessResponse{Success: true, Message: "user register successfully", UserData: models.UserData{ID: id, Name: user.Name, Email: user.Email}})

}
