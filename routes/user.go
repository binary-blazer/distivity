package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"distivity/client"
	"distivity/config/static"

	"github.com/gofiber/fiber/v2"
)

func fetchUserActivity(userID string) (map[string]interface{}, error) {
	return nil, nil
}

func UserHandler(c *fiber.Ctx) error {
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

	hostname := c.BaseURL()
	userInfo["avatar"] = fmt.Sprintf("%s/avatar/%s", hostname, userID)
	userInfo["banner"] = fmt.Sprintf("%s/banner/%s", hostname, userID)

	if bot, ok := userInfo["bot"].(bool); ok {
		userInfo["bot"] = bot
	} else {
		userInfo["bot"] = false
	}

	if system, ok := userInfo["system"].(bool); ok {
		userInfo["system"] = system
	} else {
		userInfo["system"] = false
	}

	activities, err := fetchUserActivity(userID)
	if err != nil {
		userInfo["activity_info"] = "Activity excluded because the user is not in the required Discord server"
	} else {
		userInfo["activities"] = activities
	}

	response := map[string]interface{}{
		"data":    userInfo,
		"success": true,
	}

	return c.JSON(response)
}
