package controllers

import (
	"fmt"

	"github.com/NoorUllah43/Chatapp-Backend-in-go.git/db"
	"github.com/NoorUllah43/Chatapp-Backend-in-go.git/models"
	"github.com/gofiber/fiber/v3"
)

func GetUsers(ctx fiber.Ctx) error {
	var userID int
	if ctx.Locals("userID") != nil {
		userID = ctx.Locals("userID").(int)
	}
	fmt.Println("user id in get users controller", userID)

	users, err := db.GetAllUsers()
	if err != nil {
		ctx.JSON(err)
	}



	return ctx.Status(201).JSON(models.Users{Success:true, Message: "fetched all users data", Users: users})

}
