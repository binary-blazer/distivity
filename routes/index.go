package routes

import (
	"fmt"
	"time"

	"distivity/config/static"
	"distivity/websocket"

	"github.com/gofiber/fiber/v2"
)

func IndexHandler(c *fiber.Ctx) error {
	config := static.GetConfig()
	start := time.Now()

	response := map[string]interface{}{
		"app_name": config.App.Name,
		"info":     config.App.Description,
		"code":     200,
		"success":  true,
		"authors":  config.App.Authors,
	}

	duration := time.Since(start).Milliseconds()
	response["load"] = fmt.Sprintf("%dms", duration)

	return c.JSON(response)
}

func WebsocketHandler(c *fiber.Ctx) error {
	websocket.WebsocketHandler(c)
	return nil
}
