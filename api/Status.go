package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//Status : returns status of the server
func Status(c *gin.Context) {
	c.String(http.StatusOK, "Hello from account service status module")
}
