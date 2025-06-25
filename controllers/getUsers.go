package controllers

import (

	"github.com/NoorUllah43/Chatapp-Backend-in-go.git/db"
	"github.com/NoorUllah43/Chatapp-Backend-in-go.git/models"
	"github.com/gofiber/fiber/v3"
)

func GetUsers(ctx fiber.Ctx) error {	

	users, err := db.GetAllUsers()
	if err != nil {
		ctx.JSON(err)
	}



	return ctx.Status(201).JSON(models.Users{Success:true, Message: "fetched all users data", Users: users})

}
