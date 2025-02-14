package routes

import (
	"fmt"
	"time"

	"distivity/types"

	"github.com/gofiber/fiber/v2"
)

func IndexHandler(c *fiber.Ctx, config types.Config) error {
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
