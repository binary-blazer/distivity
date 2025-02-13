package main

import (
	"distivity/config"
	"distivity/routes"
	"distivity/server"
)

func main() {
	handlers := map[string]fiber.Handler{
		"/": routes.IndexHandler,
	}
	config := config.GetConfig(handlers)
	server.Run(config, handlers)
}
