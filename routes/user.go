package routes

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"distivity/client"
	"distivity/config/static"
	"distivity/types"

	"github.com/gofiber/fiber/v2"
)

func fetchUserActivity(config types.Config, userID string) (map[string]interface{}, error) {
	session := client.GetDiscordSession()
	if session == nil || session.State == nil {
		log.Printf("Error: Discord session or state cache not found")
		return nil, fmt.Errorf("discord session or state cache not found")
	}

	activity, err := session.State.Presence(config.Discord.GuildID, userID)
	if err != nil {
		log.Printf("Error fetching user activity: %v", err)
		return nil, err
	}

	activityMap := make(map[string]interface{})
	activityBytes, err := json.Marshal(activity)
	if err != nil {
		log.Printf("Error marshalling activity: %v", err)
		return nil, err
	}
	if err := json.Unmarshal(activityBytes, &activityMap); err != nil {
		log.Printf("Error unmarshalling activity: %v", err)
		return nil, err
	}

	return activityMap, nil
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

	activities, err := fetchUserActivity(config, userID)
	if err != nil {
		userInfo["activity_info"] = fmt.Sprintf("Activity excluded because the user is not in the required Discord server: %s", config.Discord.GuildInvite)
	} else {
		userInfo["activities"] = activities["activities"]
	}

	clientStatus, ok := activities["client_status"].(map[string]interface{})
	if !ok {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to parse client status",
		})
	}

	userInfo["client_status"] = map[string]bool{
		"desktop": clientStatus["desktop"] != "",
		"mobile":  clientStatus["mobile"] != "",
		"web":     clientStatus["web"] != "",
	}

	response := map[string]interface{}{
		"data":    userInfo,
		"success": true,
	}

	return c.JSON(response)
}
