package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func Test_GetDateTime(t *testing.T) {
	request, _ := http.NewRequest(http.MethodGet, "/datetime", nil)
	response := httptest.NewRecorder()
	getDateTime(response, request)

	t.Run("testing", func(t *testing.T) {
		var responseTime time.Time
		json.NewDecoder(response.Body).Decode(&responseTime)

		got := strings.Split(responseTime.String(), " ")[0]
		want := strings.Split(time.Now().String(), " ")[0]

		if got != want {
			t.Errorf("got: %v, wanted: %v", got, want)
		}
	})
}
