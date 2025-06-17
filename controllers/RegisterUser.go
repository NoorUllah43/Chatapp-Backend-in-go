package controllers

import "github.com/gofiber/fiber/v3"

func RegisterUser(ctx fiber.Ctx) error {
	return ctx.JSON("welcome to / route")
}
