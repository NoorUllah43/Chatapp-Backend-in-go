package controllers

import (
	"github.com/NoorUllah43/Chatapp-Backend-in-go.git/db"
	"github.com/gofiber/fiber/v3"
)

func GetUsers(ctx fiber.Ctx) error {

	rows, err := db.GetAllUsers()
	if err != nil {
		ctx.JSON(err)
	}

	return ctx.JSON(rows)

}
