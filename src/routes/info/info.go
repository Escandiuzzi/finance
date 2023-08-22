package info

import (
	infoController "finance/src/controllers/info"

	"github.com/gin-gonic/gin"
)

func Info(router *gin.RouterGroup) {

	router.GET("/", infoController.GetServerStatus)
}
