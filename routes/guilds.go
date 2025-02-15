package routes

import (
	"github.com/bwmarrin/discordgo"
	"github.com/gofiber/fiber/v2"
	"distivity/client"
)

func GuildsHandler(c *fiber.Ctx) error {
	session := client.GetDiscordSession()
	if session == nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Discord session not initialized",
		})
	}

	guilds, err := session.UserGuilds(100, "", "")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch guilds",
		})
	}

	var guildList []map[string]interface{}
	for _, guild := range guilds {
		guildList = append(guildList, map[string]interface{}{
			"id":   guild.ID,
			"name": guild.Name,
		})
	}

	return c.JSON(fiber.Map{
		"guilds":  guildList,
		"success": true,
	})
}
