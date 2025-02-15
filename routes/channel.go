package routes

import (
	"github.com/bwmarrin/discordgo"
	"github.com/gofiber/fiber/v2"
	"distivity/client"
)

func ChannelHandler(c *fiber.Ctx) error {
	session := client.GetDiscordSession()
	if session == nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Discord session not initialized",
		})
	}

	channelID := c.Params("id")
	if channelID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Channel ID is required",
		})
	}

	channel, err := session.Channel(channelID)
	if err != nil {
		if err == discordgo.ErrUnauthorized {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized access to channel",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch channel information",
		})
	}

	channelInfo := map[string]interface{}{
		"id":          channel.ID,
		"name":        channel.Name,
		"type":        channel.Type,
		"guild_id":    channel.GuildID,
		"position":    channel.Position,
		"topic":       channel.Topic,
		"nsfw":        channel.NSFW,
		"last_message_id": channel.LastMessageID,
	}

	return c.JSON(fiber.Map{
		"channel": channelInfo,
		"success": true,
	})
}
