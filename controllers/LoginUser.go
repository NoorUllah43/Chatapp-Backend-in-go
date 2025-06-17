package controllers

import (
	"encoding/json"

	"github.com/NoorUllah43/Chatapp-Backend-in-go.git/models"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
)

func Login(ctx fiber.Ctx) error {

	var credentials models.LoginCredentials
	err := json.Unmarshal(ctx.Body(), &credentials)
	if err != nil {
		log.Errorf(err.Error())
	}
	return ctx.JSON("welcome ")

}
