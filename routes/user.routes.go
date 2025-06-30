package routes

import (
	"github.com/NoorUllah43/Chatapp-Backend-in-go.git/controllers"
	"github.com/NoorUllah43/Chatapp-Backend-in-go.git/middlewares"
	"github.com/gofiber/fiber/v3"
)

func UserRoutes(app *fiber.App) {

	app.Get("/user/all", middlewares.VerifyToken, controllers.GetUsers)
	app.Get("/message/history/:userId/:receiverId", middlewares.VerifyToken, controllers.GetMessageHistory)
}
