package main

import (
	"distivity/config"
	"distivity/routes"
	"distivity/server"

	"github.com/gofiber/fiber/v2"
)

func main() {
	handlers := map[string]fiber.Handler{
		"/": routes.IndexHandler,
		"/user/:id": routes.UserHandler, // Peefc
	}
	config := config.GetConfig(handlers)
	server.Run(config, handlers)
}
