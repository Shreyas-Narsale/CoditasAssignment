package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	StatusCode int                    `json:"StatusCode"`
	Message    string                 `json:"Message"`
	Data       interface{}            `json:"Data,omitempty"`
	Extras     map[string]interface{} `json:"Extras,omitempty"`
}

var statusMessages = map[int]string{
	http.StatusOK:                  "Success",
	http.StatusBadRequest:          "Bad Request",
	http.StatusUnauthorized:        "Unauthorized",
	http.StatusForbidden:           "Forbidden",
	http.StatusNotFound:            "Not Found",
	http.StatusInternalServerError: "Internal Server Error",
}

func SendResponse(c *gin.Context, statusCode int, message string, data interface{}, extras map[string]interface{}) {
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

	c.JSON(statusCode, response)
}
