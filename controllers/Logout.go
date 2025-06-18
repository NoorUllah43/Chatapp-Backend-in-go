package controllers

import (
	"time"

	"github.com/gofiber/fiber/v3"
)

func Logout(ctx fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name:     "chatappToken",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		MaxAge:   -1,
		HTTPOnly: true,
		Secure:   true,
		Path:     "/",
	}
	ctx.Cookie(&cookie)

	return ctx.JSON(fiber.Map{
		"Success": true,
		"message": "Logged out successfully",
	})
}
