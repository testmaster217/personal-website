package main

import (
	"net/http"
	"net/http/httptest"
	"slices"
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

	// When I request a binary file, it should return that data.
	t.Run("returns a requested binary file", func(t *testing.T) {
		request := NewGetReq("/testdata.dat")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		got := response.Body.Bytes()
		// "test data plz ignore"
		want := []byte{0x74, 0x65, 0x73, 0x74, 0x20, 0x64, 0x61, 0x74, 0x61, 0x20, 0x70, 0x6c, 0x7a, 0x20, 0x69, 0x67, 0x6e, 0x6f, 0x72, 0x65}

		if !slices.Equal(got, want) {
			t.Errorf("wrong page contents, got %s, want %s", got, want)
		}
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
