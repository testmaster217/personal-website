package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServePages(t *testing.T) {
	// When I request a page, it should return that page.
	t.Run("returns the requested page", func(t *testing.T) {
		request := NewGetPageReq("/testpage")
		response := httptest.NewRecorder()

		CveselServer(response, request)

		assertResponseBody(t, response.Body.String(), "test page plz ignore")
	})

	t.Run("returns a different requested page", func(t *testing.T) {
		request := NewGetPageReq("/testpage2")
		response := httptest.NewRecorder()

		CveselServer(response, request)

		assertResponseBody(t, response.Body.String(), "2nd test page plz ignore")
	})
}

func NewGetPageReq(path string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, path, nil)
	return req
}

func assertResponseBody(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("wrong page contents, got %q, want %q", got, want)
	}
}
