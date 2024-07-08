package main

import (
	"log"
	"net/http"

	date_time "github.com/codescalersinternships/datetime-server-amr/pkg/httphandler"
)

func main() {
	http.HandleFunc("/datetime", date_time.TimeDateHandler)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
