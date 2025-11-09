package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServePages(t *testing.T) {
	// When I request a page, it should return that page.
	t.Run("returns the requested page", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/testpage", nil)
		response := httptest.NewRecorder()

		// Assume this is the serevr.
		CveselServer(response, request)

		got := response.Body.String()
		want := "test page plz ignore"

		if got != want {
			t.Errorf("wrong page contents, got %q, want %q", got, want)
		}
	})

	t.Run("returns a different requested page", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/testpage2", nil)
		response := httptest.NewRecorder()

		CveselServer(response, request)

		got := response.Body.String()
		want := "2nd test page plz ignore"

		if got != want {
			t.Errorf("wrong page contents, got %q, want %q", got, want)
		}
	})
}
