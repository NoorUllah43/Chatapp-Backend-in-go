package controllers

import (
	"fmt"
	"strconv"

	"github.com/NoorUllah43/Chatapp-Backend-in-go.git/db/mongodb"
	"github.com/gofiber/fiber/v3"
)

func GetMessageHistory(c fiber.Ctx) error {
	userID := c.Params("userId")
	receiverID := c.Params("receiverId")

	if userID == "" || receiverID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "userId and receiverId are required",
		})
	}
	userId, err := strconv.Atoi(userID)
	if err != nil {
		fmt.Println("Error converting string to int:", err)
		return nil
	}

	receiverId, err := strconv.Atoi(receiverID)
	if err != nil {
		fmt.Println("Error converting string to int:", err)
		return nil
	}
	messages, err := mongodb.GetMessageHistory(userId, receiverId)
	if err != nil {
		c.JSON(fiber.Map{
			"error": err,
		})
	}

	return c.JSON(fiber.Map{
		"success":  true,
		"message":  "Message history retrieved successfully",
		"messages": messages,
	})
}
