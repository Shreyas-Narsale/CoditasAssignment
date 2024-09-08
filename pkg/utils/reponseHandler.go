package utils

import (
	"github.com/gofiber/fiber/v2"
)

type Response struct {
	StatusCode int                    `json:"StatusCode"`
	Message    string                 `json:"Message"`
	Data       interface{}            `json:"Data"`
	Extras     map[string]interface{} `json:"Extras"`
}

var statusMessages = map[int]string{
	fiber.StatusOK:                  "Success",
	fiber.StatusBadRequest:          "Bad Request",
	fiber.StatusUnauthorized:        "Unauthorized",
	fiber.StatusForbidden:           "Forbidden",
	fiber.StatusNotFound:            "Not Found",
	fiber.StatusInternalServerError: "Internal Server Error",
}

func SendResponse(c *fiber.Ctx, statusCode int, message string, data interface{}, extras map[string]interface{}) error {
	if message == "" {
		var ok bool
		message, ok = statusMessages[statusCode]
		if !ok {
			message = "Unknown Status Code"
		}
	}

	response := Response{
		StatusCode: statusCode,
		Message:    message,
		Data:       data,
		Extras:     extras,
	}

	c.Status(statusCode).JSON(response)
	return nil
}
