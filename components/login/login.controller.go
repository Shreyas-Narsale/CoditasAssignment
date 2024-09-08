package login

import (
	"SOCIAL_MEDIA_APP/pkg/config"
	"SOCIAL_MEDIA_APP/pkg/logger"
	"SOCIAL_MEDIA_APP/pkg/utils"
	PostgressQueries "SOCIAL_MEDIA_APP/platform/database/postgress/queries"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate = validator.New()

func Login(c *fiber.Ctx) error {
	logs := logger.GetLogger()
	globalConf := config.GetGlobalConfig()
	loginBody := new(LoginDto)
	if err := c.BodyParser(loginBody); err != nil {
		return utils.SendResponse(c, fiber.StatusBadRequest, "", nil, fiber.Map{
			"error":   "Cannot parse JSON",
			"details": err.Error(),
		})
	}

	if err := validate.Struct(loginBody); err != nil {
		return utils.SendResponse(c, fiber.StatusBadRequest, "", nil, fiber.Map{
			"error":   "Validation failed",
			"details": err.Error(),
		})

	}

	isExists, userId, err := PostgressQueries.CheckIfUserNameExists(loginBody.UserName)
	if err != nil {
		logs.Error().Err(err).Msg("postgress error while searching username ")
		return utils.SendResponse(c, fiber.StatusInternalServerError, "", nil, nil)
	}
	if !isExists {
		return utils.SendResponse(c, fiber.StatusBadRequest, "Invalid Username and Password", nil, nil)
	}
	logs.Debug().Msgf("UserId:%v", userId)

	salt, hashPassword, err := utils.GeneratePassword(loginBody.Password)
	if err != nil {
		logs.Error().Err(err).Msg("error while generating password")
		return utils.SendResponse(c, fiber.StatusInternalServerError, "", nil, nil)
	}
	logs.Info().Msgf("Password is: %v", loginBody.Password)
	logs.Info().Msgf("Salt is: %v", salt)
	logs.Info().Msgf("HashPassword is: %v", hashPassword)

	if !utils.VerifyPassword(loginBody.UserName, salt, hashPassword) {
		logs.Warn().Msgf("Password is InValid for user :%v", loginBody.UserName)
		return utils.SendResponse(c, fiber.StatusBadRequest, "Invalid Username and Password", nil, nil)
	}

	accessToken, refreshToken, err := GetTokens("xyz")
	if err != nil {
		logs.Error().Err(err).Msg("error while generating tokens")
		return utils.SendResponse(c, fiber.StatusInternalServerError, "", nil, nil)
	}

	// Set the refresh token in an HTTP-only cookie
	c.Cookie(&fiber.Cookie{
		Name:     "refresh_token",
		Value:    refreshToken,
		Expires:  time.Now().Add(globalConf.RefreshTokenExpTime),
		HTTPOnly: true,
	})

	return utils.SendResponse(c, fiber.StatusOK, "", fiber.Map{
		"access_token": accessToken,
	}, nil)

}
