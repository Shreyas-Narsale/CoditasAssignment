package login

import "github.com/gofiber/fiber/v2"

func Routes(router fiber.Router) {

	router.Post("/login", Login)
	//publicRouter.Post("/test")
}
