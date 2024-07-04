package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
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

		got := responseTime.Round(time.Minute)
		want := time.Now().Round(time.Minute)

		if !got.Equal(want) {
			t.Errorf("got: %v, wanted: %v", got, want)
		}
	})
}
