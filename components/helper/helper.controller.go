package helper

import (
	"SOCIAL_MEDIA_APP/pkg/utils"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate = validator.New()

func PingPong(c *fiber.Ctx) error {

	return utils.SendResponse(c, fiber.StatusOK, "", fiber.Map{
		"message": "pong",
	}, nil)

}

func TestApi(c *fiber.Ctx) error {

	p := new(TestStruct)
	if err := c.BodyParser(p); err != nil {
		return utils.SendResponse(c, fiber.StatusBadRequest, "", nil, fiber.Map{
			"error":   "Cannot parse JSON",
			"details": err.Error(),
		})
	}

	if err := validate.Struct(p); err != nil {
		return utils.SendResponse(c, fiber.StatusBadRequest, "", nil, fiber.Map{
			"error":   "Validation failed",
			"details": err.Error(),
		})

	}

	return utils.SendResponse(c, fiber.StatusOK, "", nil, nil)

}
