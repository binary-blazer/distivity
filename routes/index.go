package routes

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
)

func IndexHandler(c *fiber.Ctx) error {
	start := time.Now()

	response := map[string]interface{}{
		"app_name": "Distivity",
		"info":     "A modern RESTful API to get presence details of a discord account by ID or username",
		"code":     200,
		"success":  true,
		"authors":  []map[string]string{
			{
				"name":     "Jonas F. Franke",
				"codename": "BinaryBlazer",
				"email":    "me@binaryblazer.me",
			},
			{
				"name":     "Samuel Domke",
				"codename": "SamTheDev",
				"email":    "",
			},
		},
	}

	duration := time.Since(start).Milliseconds()
	response["load"] = fmt.Sprintf("%dms", duration)

	return c.JSON(response)
}
