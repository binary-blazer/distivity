package server

import (
	"distivity/types"
	"distivity/utils"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
)

func Run(config types.Config) {
	app := fiber.New(fiber.Config{
		Prefork:       config.Webserver.Fiber.Prefork,
		CaseSensitive: config.Webserver.Fiber.CaseSensitive,
		StrictRouting: config.Webserver.Fiber.StrictRouting,
		ServerHeader:  fmt.Sprintf("%s/%s", config.App.Name, config.App.Version),
		AppName:       fmt.Sprintf("%s - by %s", config.App.Name, utils.FormatAuthors(config.App.Authors)),
	})

	app.Static(config.Webserver.Paths.Root, config.Webserver.Paths.Static)

	app.Get("/", func(c *fiber.Ctx) error {
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
	})

	for _, route := range config.Routes {
		app.Get(route.Path, func(c *fiber.Ctx) error {
			return c.SendString(route.Handler)
		})
	}

	app.Listen(":" + fmt.Sprint(config.Webserver.Port))
}
