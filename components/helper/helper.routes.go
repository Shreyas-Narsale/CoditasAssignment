package helper

import (
	"SOCIAL_MEDIA_APP/pkg/middleware"

	"github.com/gofiber/fiber/v2"
)

func Routes(router fiber.Router) {

	router.Post("/ping", PingPong)
	router.Post("/test", middleware.JWTMiddleware, TestApi)
}
