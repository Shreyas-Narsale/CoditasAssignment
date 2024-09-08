package config

import (
	"encoding/json"
	"time"

	"github.com/gofiber/fiber/v2"
)

func GetFiberConfig() fiber.Config {

	return fiber.Config{
		ReadTimeout:                  time.Second * 30,
		JSONEncoder:                  json.Marshal,
		JSONDecoder:                  json.Unmarshal,
		CaseSensitive:                true,
		BodyLimit:                    104857600,
		StrictRouting:                true,
		ServerHeader:                 "Social Media App",
		AppName:                      "Social Media App v1.0.0",
		ReadBufferSize:               4096000,
		WriteBufferSize:              4096000,
		DisablePreParseMultipartForm: true,
		//Prefork:                      true,
	}
}
