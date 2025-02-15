package routes

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"distivity/config/static"

	"github.com/gofiber/fiber/v2"
)

func fetchUserActivity(userID string, discordToken string) (map[string]interface{}, error) {
	url := fmt.Sprintf("https://discord.com/api/v9/users/%s/activities", userID)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Printf("Error creating request: %v", err)
		return nil, err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bot %s", discordToken))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Error making request: %v", err)
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("Error response status: %v", resp.StatusCode)
		return nil, fmt.Errorf("failed to fetch user activity")
	}

	var activities map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&activities); err != nil {
		log.Printf("Error decoding response: %v", err)
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

	hostname := c.BaseURL()
	userInfo["avatar"] = fmt.Sprintf("%s/avatar/%s", hostname, userID)
	userInfo["banner"] = fmt.Sprintf("%s/banner/%s", hostname, userID)

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

	// Log the user data
	log.Printf("User data: %v", response)

	return c.JSON(response)
}
