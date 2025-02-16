package config

import (
	"distivity/types"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file")
	}
}

func GetVariables() types.Config {
	discordToken := os.Getenv("DISCORD_BOT_TOKEN")
	guildID := os.Getenv("DISCORD_GUILD_ID")
	guildInvite := os.Getenv("DISCORD_GUILD_INVITE")

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
		Discord: types.Discord{
			API: types.DiscordAPI{
				BaseURL: "https://discord.com/api/v9",
			},
			GuildID:      guildID,
			GuildInvite:  guildInvite,
			CustomStatus: "monitoring {count} users :heart:",
		},
		Routes: []types.Route{
			{
				Path: "/",
			},
			{
				Path: "/user/:id",
			},
			{
				Path: "/avatar/:id",
			},
			{
				Path: "/banner/:id",
			},
			{
				Path: "/status",
			},
		},
		Credentials: types.Credentials{
			DiscordToken: discordToken,
		},
	}
}
