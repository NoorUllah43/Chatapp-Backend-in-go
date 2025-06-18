package controllers

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/NoorUllah43/Chatapp-Backend-in-go.git/db"
	"github.com/NoorUllah43/Chatapp-Backend-in-go.git/middlewares"
	"github.com/NoorUllah43/Chatapp-Backend-in-go.git/models"
	"github.com/gofiber/fiber/v3"
)

func Login(ctx fiber.Ctx) error {

	var credentials models.LoginCredentials

	err := json.Unmarshal(ctx.Body(), &credentials)
	if err != nil {
		return ctx.JSON(models.ErrorResponse{Success: false, Message: "provide correct credentials"})
	}

	userName, storePassword, id, err := db.FindUser(credentials.Email)
	if err != nil {
		return ctx.Status(401).JSON(models.ErrorResponse{Success: false, Message: "user not found"})
	}
	// checking password is correct or not 
	checkPassword := middlewares.CheckPasswordHash(credentials.Password, storePassword)
	if !checkPassword {
		return ctx.Status(401).JSON(models.ErrorResponse{Success: false, Message: "incorrect password"})
	}

	tokenString, err := middlewares.GenerateToken(id)
	if err != nil {
		return ctx.Status(401).JSON(models.ErrorResponse{Success: false, Message: fmt.Sprintf("err in generating token %v", err)})
	}
	cookie := &fiber.Cookie{
		Name:    "chatappToken",
		Value:   tokenString,
		Expires: time.Now().Add(24 * time.Hour),
	}
	ctx.Cookie(cookie)

	return ctx.Status(201).JSON(models.SuccessResponse{Success: true, Message: "user login successfully", Data: models.UserData{ID: id, Name: userName, Email: credentials.Email}})

}
