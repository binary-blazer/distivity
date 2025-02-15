package routes

import (
	"github.com/bwmarrin/discordgo"
	"github.com/gofiber/fiber/v2"
	"distivity/client"
)

func GuildHandler(c *fiber.Ctx) error {
	session := client.GetDiscordSession()
	if session == nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Discord session not initialized",
		})
	}

	guildID := c.Params("id")
	if guildID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Guild ID is required",
		})
	}

	guild, err := session.Guild(guildID)
	if err != nil {
		if err == discordgo.ErrUnauthorized {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized access to guild",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch guild information",
		})
	}

	guildInfo := map[string]interface{}{
		"id":          guild.ID,
		"name":        guild.Name,
		"description": guild.Description,
		"owner_id":    guild.OwnerID,
		"member_count": guild.MemberCount,
	}

	return c.JSON(fiber.Map{
		"guild":   guildInfo,
		"success": true,
	})
}
