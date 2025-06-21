package server

import (
	"CODITAS_TASK/components/coditas"

	"github.com/gin-gonic/gin"
)

func LoadRoutes(router *gin.Engine) {
	apiGroup := router.Group("/v1/api")
	coditas.Routes(apiGroup)
}
