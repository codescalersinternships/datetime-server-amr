package ginhandler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// TimeDateHandler takes gin context struct and handles a request to get current date and time
func TimeDateHandler(c *gin.Context) {
	if c.Request.Method != "GET" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request"})
		return
	}
	if c.Request.URL.Path != "/datetime" {
		c.JSON(http.StatusNotFound, gin.H{"error": "Not Found"})
		return
	}
	time := time.Now().Format(time.RFC1123)
	c.String(http.StatusOK, time)
}
