package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServePages(t *testing.T) {
	server := &CveselServer{"./testpages"}

	// When I request a page, it should return that page.
	t.Run("returns the requested page", func(t *testing.T) {
		request := NewGetReq("/testpage")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertResponseBody(t, response.Body.String(), "<html><body>test page plz ignore</body></html>")
	})

	// When I request a different page, it should return that page.
	t.Run("returns a different requested page", func(t *testing.T) {
		request := NewGetReq("/testpage2")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertResponseBody(t, response.Body.String(), "<html><body>2nd test page plz ignore</body></html>")
	})

	// When I request a CSS file, it should return that file.
	t.Run("returns a requested CSS file", func(t *testing.T) {
		request := NewGetReq("/teststyles.css")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertResponseBody(t, response.Body.String(), "body {color: blue;}")
	})
}

func NewGetReq(path string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, path, nil)
	return req
}

func assertResponseBody(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("wrong page contents, got %q, want %q", got, want)
	}
}
