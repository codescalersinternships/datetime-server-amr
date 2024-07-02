package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
)

func TestTimeDateHandler(t *testing.T) {
	t.Run("Valid URL", func(t *testing.T) {
		r := gin.Default()
		r.GET("/", TimeDateHandler)

		req, err := http.NewRequest("GET", "/", nil)
		if err != nil {
			t.Fatal(err)
		}

		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)

		if status := w.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}

		expected := time.Now().Format(time.RFC1123)
		if body := w.Body.String(); body != expected {
			t.Errorf("handler returned unexpected body: got %v want %v",
				body, expected)
		}
	})

	t.Run("Invalid URL", func(t *testing.T) {
		r := gin.Default()
		r.GET("/", TimeDateHandler)

		req, err := http.NewRequest("GET", "/Invalid", nil)
		if err != nil {
			t.Fatal(err)
		}

		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)

		if status := w.Code; status != http.StatusNotFound {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusNotFound)
		}
	})
}
