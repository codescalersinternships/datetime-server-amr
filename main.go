package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func timeDateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		time := time.Now().Format(time.RFC1123)
		fmt.Fprintf(w, "Current date and time: %s", time)
	} else {
		http.Error(w, "Bad Request", http.StatusBadRequest)
	}
}

func main() {
	http.HandleFunc("/", timeDateHandler)
	log.Fatal(http.ListenAndServe(":8000", nil))
}