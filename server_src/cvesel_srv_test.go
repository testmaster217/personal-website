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

	// === HAPPY PATHS - NO HTNX === //

	// When I request a page, it should return that page.
	t.Run("returns the requested page", func(t *testing.T) {
		request := newGetReq("/testpage")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertMimeType(t, response.Result().Header.Get("Content-Type"), "text/html")
		assertResponseBody(t, response.Body.String(), "<html><body><p>test page plz ignore</p></body></html>")
	})

	// When I request a different page, it should return that page.
	t.Run("returns a different requested page", func(t *testing.T) {
		request := newGetReq("/testpage2")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertMimeType(t, response.Result().Header.Get("Content-Type"), "text/html")
		assertResponseBody(t, response.Body.String(), "<html><body><p>2nd test page plz ignore</p></body></html>")
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
		assertResponseBody(t, response.Body.String(), "<html><body><p>test collection plz ignore<br><a href=\"/testcollection/item1\">item 1</a><br><a href=\"/testcollection/item2\">item 2</a><br><a href=\"/testcollection/item3\">item 3</a></p></body></html>")
	})

	// When I visit a different page that ends with a "/" and is supposed to, it should serve that page.
	t.Run("returns a different requested collection page", func(t *testing.T)  {
		request := newGetReq("/testcollection2/")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertMimeType(t, response.Result().Header.Get("Content-Type"), "text/html")
		assertResponseBody(t, response.Body.String(), "<html><body><p>2nd test collection plz ignore<br><a href=\"/testcollection/item1\">item 1</a><br><a href=\"/testcollection/item2\">item 2</a><br><a href=\"/testcollection/item3\">item 3</a></p></body></html>")
	})

	// === HAPPY PATHS - HTMX === //

	// When I request a page, it should return that page, but not the base page content.
	t.Run("returns the requested partial page", func(t *testing.T) {
		request := newGetReq("/testpage")
		request.Header.Set("Hx-Request", "true")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertMimeType(t, response.Result().Header.Get("Content-Type"), "text/html")
		assertResponseBody(t, response.Body.String(), "<p>test page plz ignore</p>")
	})

	// When I request a non-page file, it should return that file regardless of the HTMX header.

	// === HAPPY PATHS - FALSE HTMX === //

	// When I request a page, it should return that page.
	// When I request a non-page file, it should return that file regardless of the HTMX header.

	// TODO: Add stuff involving the root page.

	// === TRAILING SLASH REDIRECTS === //

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
	t.Run("redirects missing trailing '/' to correct path", func(t *testing.T) {
		request := newGetReq("/testcollection")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertRedirect(t, response, "/testcollection/")
	})

	// When I visit a page that ends with a ".html/" and is supposed to have the "/", it should redirect to the correct path.
	t.Run("redirects trailing '.html' plus correct '/' to correct path", func(t *testing.T) {
		request := newGetReq("/testcollection.html/")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertRedirect(t, response, "/testcollection/")
	})

	// When I visit a page that ends with a ".html/" but is not supposed to have the "/", it should redirect to the correct path.
	t.Run("redirects trailing '.html' plus incorrect '/' to correct path", func(t *testing.T) {
		request := newGetReq("/testpage.html/")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertRedirect(t, response, "/testpage")
	})

	// When I visit a page that ends with a ".html" but is supposed to end with a "/", it should redirect to the correect path.
	t.Run("redirects trailing '.html' plus missing '/' to correct path", func(t *testing.T) {
		request := newGetReq("/testcollection.html")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertRedirect(t, response, "/testcollection/")
	})

	// When I visit a page that ends with a ".htm/" and is supposed to have the "/", it should redirect to the correct path.
	t.Run("redirects trailing '.htm' plus correct '/' to correct path", func(t *testing.T) {
		request := newGetReq("/testcollection.htm/")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertRedirect(t, response, "/testcollection/")
	})

	// When I visit a page that ends with a ".htm/" but is not supposed to have the "/", it should redirect to the correct path.
	t.Run("redirects trailing '.htm' plus incorrect '/' to correct path", func(t *testing.T) {
		request := newGetReq("/testpage.htm/")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertRedirect(t, response, "/testpage")
	})

	// When I visit a page that ends with a ".htm" but is supposed to end with a "/", it should redirect to the correect path.
	t.Run("redirects trailing '.htm' plus missing '/' to correct path", func(t *testing.T) {
		request := newGetReq("/testcollection.htm")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertRedirect(t, response, "/testcollection/")
	})

	// When I visit a file that isn't a page, but I have a trailing "/", redirect to the same path without it.
	t.Run("redirects trailing '/' on non-page file to correct path", func(t *testing.T)  {
		request := newGetReq("/teststyles.css/")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertRedirect(t, response, "/teststyles.css")
	})

	// TODO: Add stuff related to the root.

	// === 500 ERRORS === //

	// TODO: When I try to request a page that exists, but I can't get it's metadata, the server should return a 500 error.
	// TODO: When I try to request a file that exists, but I can't get it's metadata, the server should return a 500 error.
	// TODO: When I try to request a collection page that exists, but I can't get the directory's metadata, the server should return a 500 error.
	// When I try to request a file, but that file isn't a regular file (it's instead a symlink or something OR it's a directory, but doesn't have a corresponding HTML file), the server should return a 500 error.
	// When I try to request a page, but that page isn't a regular file (it's instead a directory or a symlink or something), the server should return a 500 error.
	// When I try to request a page, but that page isn't an HTML file (it says it is, but it actually contains other content), the server should return a 500 error.
	// When I try to request a collection page, but the directory is not a directory (it's a symlink or something), the server should return a 500 error.

	// === 404 ERRORS === //

	// When I request a file that is not a page, but that file (the path with no modifications) does not exist, the server should return a 404 error.
	// When I request a page, but that page (the path I requested wirth ".html" added to the end) does not exist, the server should return a 404 error.
	// TODO: Decide if I want versions of these tests with different combinations of trailing "/", ".htm", and ".html".

	// === 405 ERRORS === //

	// When I try to use a POST method, the server should return a 405 error.
	// When I try to use a PUT method, the server should return a 405 error.
	// When I try to use a PATCH method, the server should return a 405 error.
	// When I try to use a DELETE method, the server should return a 405 error.
	// When I try to use a CONNECT method, the server should return a 405 error. (This may change in the future.)
	// When I try to use a TRACE method, the server should return a 405 error. (This may change in the future.)
	// When I try to use a nonstandard method, the server should return a 405 error.

	// === OTHER ERRORS === //

	// When a path maps to a file outside the server's docRoot, the server should return a 403 error.
	// When I try to request the base page ("__base__.html"), the server should return a 403 error.
	
	// === OTHER TESTS === //

	// When I try to use an OPTIONS method, the server should return a 204 response with an Allow header listing all allowed methods. (Currently, these are OPTIONS, GET, and HEAD for the entire server, but this may change in the future.)
	// When I try to use a HEAD method, the server should return the same response that it would send for a GET method, but without a response body.
	// TODO: Add HTTPS stuff.
	// TODO: Find out what standards this server needs to comply with and make sure it does (unless this would cause security issues).
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
