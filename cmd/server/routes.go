package server

import (
	"SOCIAL_MEDIA_APP/components/helper"
	"SOCIAL_MEDIA_APP/components/login"

	"github.com/gofiber/fiber/v2"
)

func LoadRoutes(app *fiber.App) {
	router := app.Group("/portal/api")

	login.Routes(router)
	helper.Routes(router)
}
