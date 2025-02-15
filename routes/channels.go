package routes

import (
	"github.com/bwmarrin/discordgo"
	"github.com/gofiber/fiber/v2"
	"distivity/client"
)

func ChannelsHandler(c *fiber.Ctx) error {
	session := client.GetDiscordSession()
	if session == nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Discord session not initialized",
		})
	}

	guildID := c.Query("guild_id")
	if guildID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Guild ID is required",
		})
	}

	channels, err := session.GuildChannels(guildID)
	if err != nil {
		if err == discordgo.ErrUnauthorized {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized access to guild channels",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch guild channels",
		})
	}

	var channelList []map[string]interface{}
	for _, channel := range channels {
		channelList = append(channelList, map[string]interface{}{
			"id":   channel.ID,
			"name": channel.Name,
			"type": channel.Type,
		})
	}

	return c.JSON(fiber.Map{
		"channels": channelList,
		"success":  true,
	})
}
