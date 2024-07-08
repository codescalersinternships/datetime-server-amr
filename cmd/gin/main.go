package main

import (
	"log"

	"github.com/gin-gonic/gin"

	date_time "github.com/codescalersinternships/datetime-server-amr/pkg/ginhandler"
)

func main() {
	r := gin.Default()
	r.HandleMethodNotAllowed = true
	r.GET("/datetime", date_time.TimeDateHandler)
	log.Fatal(r.Run(":8000"))
}
