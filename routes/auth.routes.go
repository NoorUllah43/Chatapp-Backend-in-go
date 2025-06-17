package routes

import (
	"github.com/NoorUllah43/Chatapp-Backend-in-go.git/controllers"
	"github.com/gofiber/fiber/v3"
)

func AuthRoutes(app *fiber.App) {

	// api endpoints
	app.Get("/", func(ctx fiber.Ctx) error {
		return ctx.JSON("welcome to / route")
	})

	app.Post("/auth/login", controllers.Login)
	app.Post("/auth/registerUser", controllers.RegisterUser)
	

}
