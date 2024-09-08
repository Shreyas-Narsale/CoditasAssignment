package middleware

import (
	"SOCIAL_MEDIA_APP/pkg/config"
	"SOCIAL_MEDIA_APP/pkg/jwt"
	"SOCIAL_MEDIA_APP/pkg/logger"
	"SOCIAL_MEDIA_APP/pkg/utils"
	"strings"
	"sync"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	Logger "github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/rs/zerolog"
)

var (
	once       sync.Once
	logs       zerolog.Logger
	globalConf *config.GlobalConfig
)

func configureGlobals() {
	once.Do(func() {
		logs = logger.GetLogger()
		globalConf = config.GetGlobalConfig()
	})
}

func FiberMiddleware(app *fiber.App) {
	// Logger middleware
	app.Use(Logger.New())
	// CORS middleware
	app.Use(cors.New())
}

// JWTMiddleware is the middleware function for verifying JWT tokens
func JWTMiddleware(c *fiber.Ctx) error {
	configureGlobals()
	jwtSecret := globalConf.JwtSecretKeyAccessToken

	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return utils.SendResponse(c, fiber.StatusUnauthorized, "", nil, fiber.Map{
			"error":   "Unauthorized",
			"details": "Missing Authorization header",
		})
	}

	if !strings.HasPrefix(authHeader, "Bearer ") {
		return utils.SendResponse(c, fiber.StatusUnauthorized, "", nil, fiber.Map{
			"error":   "Unauthorized",
			"details": "Invalid Authorization header format",
		})
	}

	tokenString := strings.TrimPrefix(authHeader, "Bearer ")

	userID, err := jwt.VerifyJWT(tokenString, jwtSecret)
	if err != nil {
		logs.Error().Err(err).Msg("verifyToken error:")
		return utils.SendResponse(c, fiber.StatusUnauthorized, "", nil, fiber.Map{
			"error":   "Unauthorized",
			"details": "Invalid Token",
		})
	}

	c.Locals("userID", userID)

	return c.Next()
}
