package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func TestServer(t *testing.T) {
	// TODO: Make the server work in such a way that it can have a different
	// document root folder in prod than when running these tests.

	tests := []struct {
		// Parameters
		method, htmxHeader, reqPath string
		// Wanted results
		wantStatusCode       int
		wantPageContentsPath string
	}{
		// Test cases go here.
		// - Happy path HTMX.
		// - Happy path no HTMX.
		{http.MethodGet, "", "bio.html", http.StatusOK, "bio.html"},
		// - Happy path false HTMX.
		// - Happy path non-HTML.
		// - Happy path root HTMX.
		// - Happy path root no HTMX.
		// - Happy path root false HTMX.
		// - HEAD request.
		// - Illegal method.
		// - Page does not exist.
		// - Page is not a file.
		// - Illegal page.
		// - Path ends in "/".
		// - Path ends in ".html".
		// - Path ends in ".htm".
		// - Path is "/index.html".
		// - Path is "/index.htm".
		// - Path is "/index".
		// - Path is invalid.
	}

	// Run the tests.
	var sb strings.Builder
	for _, tt := range tests {
		testname := fmt.Sprintf("%v %v HTMX: %q", tt.method, tt.reqPath, tt.htmxHeader)
		t.Run(testname, func(t *testing.T) {
			// Try to read the file with the desired response contents.
			// If this fails, the test is not runnable.
			sb.Reset()
			sb.WriteString("test_files/")
			sb.WriteString(tt.wantPageContentsPath)
			wantedResBody, err := os.ReadFile(sb.String())
			if err != nil {
				t.Fatalf("failed to run test, %v", err)
			}
			// Setup.
			req := httptest.NewRequest(tt.method, tt.reqPath, nil)
			w := httptest.NewRecorder()
			// If we're testing an HTMX request, add the necessary header.
			if tt.htmxHeader != "" {
				req.Header.Add("HX-Request", tt.htmxHeader)
			}

			// Make the request and get the response.
			handler(w, req)

			// Compare the results with the desired ones.
			res := w.Result()
			// Check status code.
			if res.StatusCode != tt.wantStatusCode {
				t.Errorf("status code was %d, want %d", res.StatusCode, tt.wantStatusCode)
			}
			// Check body.
			responseBody, err := io.ReadAll(res.Body)
			if err != nil {
				t.Errorf("%v", err)
			}
			if bytes.Equal(responseBody, wantedResBody) {
				t.Errorf("response body was %q, want %q", responseBody, wantedResBody)
			}
		})
	}
}
