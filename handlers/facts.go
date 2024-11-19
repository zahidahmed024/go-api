package handlers

import "github.com/gofiber/fiber/v3"

func Home(c fiber.Ctx) error {
	// Send a string response to the client
	return c.SendString("dont know what to do")
}
