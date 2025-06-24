package main

import (
	"fmt"
	"log"

	"github.com/NoorUllah43/Chatapp-Backend-in-go.git/db"
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
	db.ConnectPostgresql()
	// db.ConnectMongoDB()
	
	app.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:5173"},
		AllowHeaders: []string{"Origin", "Content-Type", "Accept", "Authorization"},
		AllowMethods: []string{"GET", "POST", "HEAD", "PUT", "DELETE", "PATCH", "OPTIONS"},
		AllowCredentials: true,
	}))

	routes.AuthRoutes(app)
	routes.UserRoutes(app)

	log.Fatal(app.Listen(":5000"))
	


}
