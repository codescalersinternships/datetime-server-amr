package main

import (
	"log"
	"net/http"

	date_time "github.com/codescalersinternships/datetime-server-amr/http/handler"
)

func main() {
	http.HandleFunc("/", date_time.TimeDateHandler)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
