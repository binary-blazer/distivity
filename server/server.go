package server

import (
	"distivity/types"
	"distivity/utils"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func Run(config types.Config, handlers map[string]fiber.Handler) {
	app := fiber.New(fiber.Config{
		Prefork:       config.Webserver.Fiber.Prefork,
		CaseSensitive: config.Webserver.Fiber.CaseSensitive,
		StrictRouting: config.Webserver.Fiber.StrictRouting,
		ServerHeader:  fmt.Sprintf("%s/%s", config.App.Name, config.App.Version),
		AppName:       fmt.Sprintf("%s - by %s", config.App.Name, utils.FormatAuthors(config.App.Authors)),
	})

	app.Static(config.Webserver.Paths.Root, config.Webserver.Paths.Static)

	for path, handler := range handlers {
		app.Get(path, func(c *fiber.Ctx) error {
			return handler(c)
		})
	}

	app.Listen(":" + fmt.Sprint(config.Webserver.Port))
}
