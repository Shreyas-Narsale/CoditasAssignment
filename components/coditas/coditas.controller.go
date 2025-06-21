package coditas

import (
	"CODITAS_TASK/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PingPong(c *gin.Context) {
	utils.SendResponse(c, http.StatusOK, "", map[string]string{
		"message": "pong",
	}, nil)
}

func TestApi(c *gin.Context) {
	var userBody UserDto
	if err := c.ShouldBindJSON(&userBody); err != nil {
		utils.SendResponse(c, http.StatusBadRequest, "", nil, map[string]interface{}{
			"error":   "Cannot parse JSON",
			"details": err.Error(),
		})
		return
	}

	if validationErrors := utils.ValidateStruct(userBody); validationErrors != nil {
		utils.SendResponse(c, http.StatusBadRequest, "", nil, map[string]interface{}{
			"error":   "Validation failed",
			"details": validationErrors,
		})
		return
	}

	utils.SendResponse(c, http.StatusOK, "", map[string]interface{}{
		"message": "User Validated Successfully",
	}, nil)
}
