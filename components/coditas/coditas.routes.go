package coditas

import (
	"github.com/gin-gonic/gin"
)

func Routes(router *gin.RouterGroup) {
	router.POST("/ping", PingPong)
	router.POST("/test", TestApi)
}
