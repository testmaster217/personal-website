package main

import (
	"net/http"
	"net/http/httptest"
	"slices"
	"strings"
	"testing"
)

func TestServePages(t *testing.T) {
	server := NewCveselServer("./testpages")

	// When I request a page, it should return that page.
	t.Run("returns the requested page", func(t *testing.T) {
		request := NewGetReq("/testpage")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		AssertMimeType(t, response.Result().Header.Get("Content-Type"), "text/html")
		assertResponseBody(t, response.Body.String(), "<html><body>test page plz ignore</body></html>")
	})

	// When I request a different page, it should return that page.
	t.Run("returns a different requested page", func(t *testing.T) {
		request := NewGetReq("/testpage2")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		AssertMimeType(t, response.Result().Header.Get("Content-Type"), "text/html")
		assertResponseBody(t, response.Body.String(), "<html><body>2nd test page plz ignore</body></html>")
	})

	// When I request a CSS file, it should return that file.
	t.Run("returns a requested CSS file", func(t *testing.T) {
		request := NewGetReq("/teststyles.css")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		AssertMimeType(t, response.Result().Header.Get("Content-Type"), "text/css")
		assertResponseBody(t, response.Body.String(), "body {color: blue;}")
	})

	// When I request a binary file, it should return that data.
	t.Run("returns a requested binary file", func(t *testing.T) {
		request := NewGetReq("/testdata.dat")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		AssertMimeType(t, response.Result().Header.Get("Content-Type"), "application/octet-stream")

		got := response.Body.Bytes()
		// "test data plz ignore" plus a bunch of random bytes
		want := []byte{
			0x74, 0x65, 0x73, 0x74, 0x20, 0x64, 0x61, 0x74,
			0x61, 0x20, 0x70, 0x6c, 0x7a, 0x20, 0x69, 0x67,
			0x6e, 0x6f, 0x72, 0x65, 0xBA, 0x08, 0x72, 0x2d,
			0xfc, 0x7a, 0x21, 0x4e, 0xd2, 0xff, 0x00, 0x8c,
		}

		if !slices.Equal(got, want) {
			t.Errorf("wrong page contents, got %s, want %s", got, want)
		}
	})
}

func NewGetReq(path string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, path, nil)
	return req
}

func AssertMimeType(t testing.TB, got, want string) {
	t.Helper()
	if !strings.Contains(got, want) {
		t.Errorf("content is the wrong MIME type, got %q, want %q", got, want)
	}
}

func assertResponseBody(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("wrong page contents, got %q, want %q", got, want)
	}
}
