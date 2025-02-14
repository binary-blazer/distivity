package main

import (
	"distivity/config"
	"distivity/routes"
	"distivity/server"

	"github.com/gofiber/fiber/v2"
)

func main() {
	handlers := map[string]fiber.Handler{
		"/": func(c *fiber.Ctx) error {
			return routes.IndexHandler(c, config)
		},
	}
	config := config.GetConfig(handlers)
	server.Run(config, handlers)
}
