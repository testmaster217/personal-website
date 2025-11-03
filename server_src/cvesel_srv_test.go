package main

import (
	"fmt"
	"io"
	"net/http/httptest"
	"testing"
)

func TestServer(t *testing.T) {
	tests := []struct {
		params struct {
			method, htmxHeader, reqPath string
		}
		want struct {
			statusCode int
			resBody    string
		}
	}{
		// Test cases go here.
		// - Happy path HTMX.
		// - Happy path no HTMX.
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

	// Global Setup.
	w := httptest.NewRecorder()

	// Run the tests.
	for _, tt := range tests {
		testname := fmt.Sprintf("%v %v HTMX: %q", tt.params.method, tt.params.reqPath, tt.params.htmxHeader)
		t.Run(testname, func(t *testing.T) {
			// Test-specific setup.
			req := httptest.NewRequest(tt.params.method, tt.params.reqPath, nil)
			// If we're testing an HTMX request, add the necessary header.
			if tt.params.htmxHeader != "" {
				req.Header.Add("HX-Request", tt.params.htmxHeader)
			}

			// Make the request and get the response.
			handler(w, req)

			// Compare the results with the desired ones.
			res := w.Result()
			// Check status code.
			if res.StatusCode != tt.want.statusCode {
				t.Errorf("status code was %d, want %d", res.StatusCode, tt.want.statusCode)
			}
			// Check body.
			responseBody, err := io.ReadAll(res.Body)
			if err != nil {
				t.Errorf("%v", err)
			}
			stringifiedResBody := string(responseBody)
			if stringifiedResBody != tt.want.resBody {
				t.Errorf("response body was %q, want %q", stringifiedResBody, tt.want.resBody)
			}
		})
	}
}
