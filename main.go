package main

import (
	"distivity/client"
	"distivity/config/module"
	"distivity/routes"
	"distivity/server"
	"distivity/websocket" // P07de

	"github.com/gofiber/fiber/v2"
)

func main() {
	client.InitDiscordClient()

	handlers := map[string]fiber.Handler{
		"/":           routes.IndexHandler,
		"/user/:id":   routes.UserHandler,
		"/avatar/:id": routes.UserAvatarHandler,
		"/banner/:id": routes.UserBannerHandler,
		"/status":     routes.StatusHandler,
		"/ws":         websocket.WebsocketHandler, // P77d6
	}
	config := module.GetConfig(handlers)
	server.Run(config, handlers)
}
