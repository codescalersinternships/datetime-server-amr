package main

import (
	"log"
	"net/http"
	"github.com/codescalersinternships/datetime-server-amr/handler"
)

func main() {
	http.HandleFunc("/", timeDateHandler)
	log.Fatal(http.ListenAndServe(":8000", nil))
}