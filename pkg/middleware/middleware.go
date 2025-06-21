package middleware

import (
	"CODITAS_TASK/pkg/logger"
	"time"

	"github.com/gin-gonic/gin"
)

func GinMiddleware(router *gin.Engine) {
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(APILatencyLogger())
}

func APILatencyLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		logs := logger.GetLogger()
		start := time.Now()
		c.Next()
		latency := time.Since(start)
		status := c.Writer.Status()

		logs.Info().
			Str("method", c.Request.Method).
			Str("path", c.Request.URL.Path).
			Int("status", status).
			Dur("latency", latency).
			Msg("API request completed")
	}
}
