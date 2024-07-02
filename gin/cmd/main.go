package main

import (
	"log"

	"github.com/gin-gonic/gin"

	date_time "github.com/codescalersinternships/datetime-server-amr/gin/handler"
)

func main() {
	r := gin.Default()
	r.GET("/", date_time.TimeDateHandler)
	log.Fatal(r.Run(":8000"))
}
