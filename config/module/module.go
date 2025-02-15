package module

import (
	"distivity/config"
	"distivity/types"

	"github.com/gofiber/fiber/v2"
)

func GetConfig(handlers map[string]fiber.Handler) types.Config {
	configVariables := config.GetVariables()
	configVariables.Handlers = handlers
	return configVariables
}
