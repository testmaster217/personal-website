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

	// === HAPPY PATHS === //

	// When I request a page, it should return that page.
	t.Run("returns the requested page", func(t *testing.T) {
		request := newGetReq("/testpage")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertMimeType(t, response.Result().Header.Get("Content-Type"), "text/html")
		assertResponseBody(t, response.Body.String(), "<html><body>test page plz ignore</body></html>")
	})

	// When I request a different page, it should return that page.
	t.Run("returns a different requested page", func(t *testing.T) {
		request := newGetReq("/testpage2")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertMimeType(t, response.Result().Header.Get("Content-Type"), "text/html")
		assertResponseBody(t, response.Body.String(), "<html><body>2nd test page plz ignore</body></html>")
	})

	// When I request a CSS file, it should return that file.
	t.Run("returns a requested CSS file", func(t *testing.T) {
		request := newGetReq("/teststyles.css")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertMimeType(t, response.Result().Header.Get("Content-Type"), "text/css")
		assertResponseBody(t, response.Body.String(), "body {color: blue;}")
	})

	// When I request a binary file, it should return that data.
	t.Run("returns a requested binary file", func(t *testing.T) {
		request := newGetReq("/testdata.dat")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertMimeType(t, response.Result().Header.Get("Content-Type"), "application/octet-stream")

		got := response.Body.Bytes()
		// "test data plz ignore" followed by a bunch of random bytes
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

	// When I visit a page that ends with a "/" and is supposed to, it should serve that page.
	t.Run("returns a requested collection page", func(t *testing.T)  {
		request := newGetReq("/testcollection/")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertMimeType(t, response.Result().Header.Get("Content-Type"), "text/html")
		assertResponseBody(t, response.Body.String(), "<html><body>test collection plz ignore<br><a href=\"/testcollection/item1\">item 1</a><br><a href=\"/testcollection/item2\">item 2</a><br><a href=\"/testcollection/item3\">item 3</a></body></html>")
	})

	// When I visit a different page that ends with a "/" and is supposed to, it should serve that page.
	t.Run("returns a different requested collection page", func(t *testing.T)  {
		request := newGetReq("/testcollection2/")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertMimeType(t, response.Result().Header.Get("Content-Type"), "text/html")
		assertResponseBody(t, response.Body.String(), "<html><body>2nd test collection plz ignore<br><a href=\"/testcollection/item1\">item 1</a><br><a href=\"/testcollection/item2\">item 2</a><br><a href=\"/testcollection/item3\">item 3</a></body></html>")
	})

	// === REDIRECTS === //

	// When I try to visit a path that ends with ".html", redirect to the same path without the trailing extension.
	t.Run("redirects trailing '.html' to correct path", func(t *testing.T)  {
		request := newGetReq("/testpage.html")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertRedirect(t, response, "/testpage")
	})

	// When I try to visit a different path that ends with ".html", redirect to that path without the trailing extension.
	t.Run("redirects different trailing '.html' to correct path", func(t *testing.T)  {
		request := newGetReq("/testpage2.html")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertRedirect(t, response, "/testpage2")
	})

	// When I try to visit a path that ends with ".htm", redirect to the same path without the trailing extension.
	t.Run("redirects trailing '.htm' to correct path", func(t *testing.T) {
		request := newGetReq("/testpage.htm")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertRedirect(t, response, "/testpage")
	})

	// When I try to visit a different path that ends with ".htm", redirect to that path without the trailing extension.
	t.Run("redirects different trailing '.htm' to correct path", func(t *testing.T)  {
		request := newGetReq("/testpage2.htm")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertRedirect(t, response, "/testpage2")
	})

	// When I visit a page that ends with a "/" but is not supposed to, it should redirect to the correct path.
	t.Run("redirects incorrect trailing '/' to correct path", func(t *testing.T) {
		request := newGetReq("/testpage/")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertRedirect(t, response, "/testpage")
	})
	// When I visit a page that does not end with a "/" but is supposed to, it should redirect to the correect path.
	// When I visit a page that ends with a ".html/" and is supposed to have the "/", it should redirect to the correct path.
	// When I visit a page that ends with a ".html/" but is not supposed to have the "/", it should redirect to the correct path.
	// When I visit a page that ends with a ".html" but is supposed to end with a "/", it should redirect to the correect path.
	// When I visit a page that ends with a ".htm/" and is supposed to have the "/", it should redirect to the correct path.
	// When I visit a page that ends with a ".htm/" but is not supposed to have the "/", it should redirect to the correct path.
	// When I visit a page that ends with a ".htm" but is supposed to end with a "/", it should redirect to the correect path.
}

func newGetReq(path string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, path, nil)
	return req
}

func assertMimeType(t testing.TB, got, want string) {
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

func assertRedirect(t testing.TB, got *httptest.ResponseRecorder, wantRedirectUrl string) {
	t.Helper()
	// Don't need to actually follow the redirect, just need to know that we got the right response.
	gotStatus := got.Result().StatusCode
	wantStatus := http.StatusMovedPermanently

	if gotStatus != wantStatus {
		t.Errorf("wrong HTTP response, got %d, want %d", gotStatus, wantStatus)
	}

	// Do also want to make sure that we're being redirected to the right place, though.
	gotRedirectUrl := got.Result().Header.Get("Location")

	if gotRedirectUrl != wantRedirectUrl {
		t.Errorf("wrong redirect URL, got %q, want %q", gotRedirectUrl, wantRedirectUrl)
	}

	// Also need to make sure the body is correct.
	// (This is partly a side effect of my TDD-ing, and partly a response to
	// a security breach I found out about on the 0.2 version of my website
	// which wasn't using this server.)
	assertResponseBody(t, got.Body.String(), "<a href=\"" + wantRedirectUrl + "\">Moved Permanently</a>.\n\n")
}
