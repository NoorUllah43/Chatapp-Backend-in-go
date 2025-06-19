package routes

import (
	"github.com/NoorUllah43/Chatapp-Backend-in-go.git/controllers"
	"github.com/gofiber/fiber/v3"
)

func UserRoutes(app *fiber.App) {

	app.Get("/user/all", controllers.GetUsers)
}