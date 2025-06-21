package server

import (
	"CODITAS_TASK/components/coditas"
	"CODITAS_TASK/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	middleware.GinMiddleware(router)
	LoadRoutes(router)
	return router
}

func LoadRoutes(router *gin.Engine) {
	apiGroup := router.Group("/v1/api")
	coditas.Routes(apiGroup)
}
