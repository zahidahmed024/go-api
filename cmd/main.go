package main

import (
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/zahidahmed024/go-api/database"
)

func main() {
	database.ConnectDb()
	// Initialize a new Fiber app
	app := fiber.New()
	setupRoutes(app)
	// Define a route for the GET method on the root path '/'
	// app.Get("/", func(c fiber.Ctx) error {
	// 	// Send a string response to the client
	// 	return c.SendString("hi zahid ahmed dur")
	// })

	// Start the server on port 3000
	log.Fatal(app.Listen(":3000"))
}
