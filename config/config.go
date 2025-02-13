package config

import (
	"distivity/routes"
	"distivity/types"
)

func GetConfig() types.Config {
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
			Description: "A modern RESTful API to get presence details of a discord account by ID or username",
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
				Path:    "/",
				Handler: routes.IndexHandler,
			},
		},
	}
}
