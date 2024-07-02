package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func TimeDateHandler(c *gin.Context) {
	if c.Request.Method != "GET" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request"})
		return
	}
	if c.Request.URL.Path != "/" {
		c.JSON(http.StatusNotFound, gin.H{"error": "Not Found"})
		return
	}
	time := time.Now().Format(time.RFC1123)
	c.String(http.StatusOK, time)
}
