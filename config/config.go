package config

import (
	"distivity/types"

	"github.com/gofiber/fiber/v2"
)

func GetConfig(handlers map[string]fiber.Handler) types.Config {
	return types.Config{
		Webserver: types.Webserver{
			Port: 3000,
			Paths: types.WebserverPaths{
				Root:   "/",
				Static: "./public",
			},
			Fiber: types.FiberSettings{
				Prefork:       false,
				CaseSensitive: true,
				StrictRouting: true,
			},
		},
		App: types.App{
			Name:        "Distivity",
			Description: "A modern RESTful API to get presence details of a discord account by ID",
			Version:     "1.0.0",
			Environment: "development",
			Authors: []types.Author{
				{
					Name:     "Jonas F. Franke",
					Codename: "BinaryBlazer",
					Email:    "me@binaryblazer.me",
				},
				{
					Name:     "Samuel Domke",
					Codename: "SamTheDev",
					Email:    "",
				},
			},
		},
		Routes: []types.Route{
			{
				Path: "/",
			},
		},
		Handlers: handlers,
	}
}
