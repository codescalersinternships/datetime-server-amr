package handler

import (
	"fmt"
	"net/http"
	"time"
)

func timeDateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	time := time.Now().Format(time.RFC1123)
	fmt.Fprint(w, time)
}
