package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func Test_GetDateTime(t *testing.T) {
	router := gin.Default()
	router.GET("/datetime", getDateTime)

	request, _ := http.NewRequest(http.MethodGet, "/datetime", nil)
	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)

	t.Run("testing", func(t *testing.T) {
		var responseJSON map[string]string
		json.NewDecoder(response.Body).Decode(&responseJSON)
		responseTimeStr, _ := responseJSON["datetime"]
		responseTime, _ := time.Parse(time.RFC3339, responseTimeStr)
		got := strings.Split(responseTime.String(), " ")[0]
		want := strings.Split(time.Now().String(), " ")[0]

		if got != want {
			t.Errorf("got: %v, wanted: %v", got, want)
		}
	})
}
