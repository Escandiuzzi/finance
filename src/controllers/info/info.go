package info

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetServerStatus(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "Server is up and running :)")
}
