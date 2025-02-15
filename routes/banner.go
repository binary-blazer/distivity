package routes

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"distivity/config/static"

	"github.com/gofiber/fiber/v2"
)

func UserBannerHandler(c *fiber.Ctx) error {
	config := static.GetConfig()

	userID := c.Params("id")
	if userID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "User ID is required",
		})
	}

	if config.Credentials.DiscordToken == "" {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Discord bot token is not set",
		})
	}

	url := fmt.Sprintf("%s/users/%s", config.Discord.API.BaseURL, userID)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create request",
		})
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bot %s", config.Credentials.DiscordToken))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch user information",
		})
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return c.Status(resp.StatusCode).JSON(fiber.Map{
			"error": "Failed to fetch user information",
		})
	}

	var userInfo map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&userInfo); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to decode user information",
		})
	}

	bannerID, ok := userInfo["banner"].(string)
	if !ok || bannerID == "" {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Banner not found, user may not have Nitro",
		})
	}

	bannerURL := fmt.Sprintf("https://cdn.discordapp.com/banners/%s/%s.png", userID, bannerID)
	bannerResp, err := http.Get(bannerURL)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch banner",
		})
	}
	defer bannerResp.Body.Close()

	if bannerResp.StatusCode != http.StatusOK {
		return c.Status(bannerResp.StatusCode).JSON(fiber.Map{
			"error": "Failed to fetch banner",
		})
	}

	c.Set("Content-Type", bannerResp.Header.Get("Content-Type"))
	c.Set("Content-Length", bannerResp.Header.Get("Content-Length"))

	_, err = io.Copy(c.Response().BodyWriter(), bannerResp.Body)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to send banner",
		})
	}

	return nil
}
