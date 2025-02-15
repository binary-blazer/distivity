package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"distivity/config/static"

	"github.com/gofiber/fiber/v2"
)

func fetchUserActivity(userID string, discordToken string) (map[string]interface{}, error) {
	url := fmt.Sprintf("https://discord.com/api/v9/users/%s/activities", userID)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bot %s", discordToken))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch user activity")
	}

	var activities map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&activities); err != nil {
		return nil, err
	}

	return activities, nil
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

	url := fmt.Sprintf("https://discord.com/api/v9/users/%s", userID)
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

	activities, err := fetchUserActivity(userID, config.Credentials.DiscordToken)
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
