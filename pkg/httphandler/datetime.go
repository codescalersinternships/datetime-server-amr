package httphandler

import (
	"fmt"
	"net/http"
	"time"
)

func TimeDateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	if r.URL.Path != "/datetime" {
		http.NotFound(w, r)
		return
	}
	time := time.Now().Format(time.RFC1123)
	fmt.Fprint(w, time)
}
