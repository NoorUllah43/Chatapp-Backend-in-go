package main

import (
	"fmt"
	"log"

	"github.com/NoorUllah43/Chatapp-Backend-in-go.git/routes"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {

	app := fiber.New()

	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("env err:", err)
	}

	app.Use(cors.Config{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{"Origin", "Content-Type", "Accept", "Authorization"},
		AllowMethods: []string{"GET", "POST", "HEAD", "PUT", "DELETE", "PATCH", "OPTIONS"},
	})


	routes.AuthRoutes(app)


	log.Fatal(app.Listen(":3000"))

}
