package main

import (
	"github.com/gofiber/fiber/v3"
	"github.com/zahidahmed024/go-api/handlers"
)

func setupRoutes(app *fiber.App) {
	app.Get("/", handlers.Home)
}
