package controllers

import "github.com/gofiber/fiber/v3"

func Logout(ctx fiber.Ctx) error {
	ctx.ClearCookie("chatappToken")

	return ctx.SendString("Cookie cleared successfully!")
}
