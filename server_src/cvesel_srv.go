package main

import (
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
)

var isFileRegexp = regexp.MustCompile(`\.\w+/?$`)

type CveselServer struct {
	// The base file server that my custom server will extend.
	// Should not be set manually, will be set by the NewCveselServer function.
	baseServer http.Handler
	// The document root for the server.
	docRoot string
}

func NewCveselServer(docRoot string) *CveselServer {
	fileServer := http.FileServer(http.Dir(docRoot))
	return &CveselServer{fileServer, docRoot}
}

// Should do different things depending on what is requested:
// - If the path ends with a "." followed by anything other than "html" or "htm",
// - - THIS IS CUT: If the file is executable and is not a JS file, it should be
// executed, and the results should be served.
// - - Otherwise, the file should be served as is.
// - TODO: If the requested path is "/", or "/index[.htm|.html]" (possibly with
// a "/" at the end), redirect to "/" if necessary, and serve the contents of
// "index.html".
// - PARTIAL TODO; HANDLE TRAILING SLASH: If this does not apply, and the path
// ends with ".html" or ".htm" (not counting a trailing "/"), redirect to the
// equivalent path without the extension.
// - PARTIAL TODO; HANDLE TRAILING SLASH: If the path does not end with a file
// extension (not counting a trailing "/"), serve the file that has that name
// with ".html" added onto the end.
// - TODO: "Collection" pages (pages whose purpose is to give links to other
// pages, i.e. "collinvesel.me/blog/"; these will have the same name as a folder
// in the same directory) should have a trailing "/" in their paths, but other
// pages should not. Redirect as necessary.
// Whatever gets served after the above are evaluated should be served
// differently depending on whether the request came from HTMX or not, as
// indicated by the "HX-Request" header:
// - If the file is NOT an HTML file, serve it as is.
// - PARTIAL TODO; DOESN'T WORK WITH PARTIAL PAGES YET: If the header is msising
// or set to "false", combine the data to be served with a base page and serve
// that.
// - TODO: Otherwise, serve the data as is.
//
// TODO: Find out what else I need to make sure this server does, what other
// standards it needs to comply with, what headers the responses need to have,
// what the values of those headers should be, etc.
func (srv *CveselServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	pathWithoutTrailingSlash, hasTrailingSlash := strings.CutSuffix(r.URL.Path, "/")

	// === Redirect logic === //
	redirectPath := ""
	// If the request path ends with ".html", redirect to the same path without it.
	if strings.HasSuffix(pathWithoutTrailingSlash, ".html") {
		redirectPath = pathWithoutTrailingSlash[:len(pathWithoutTrailingSlash)-len(".html")]
	}
	// Ditto for trailing ".htm"
	if strings.HasSuffix(pathWithoutTrailingSlash, ".htm") {
		redirectPath = pathWithoutTrailingSlash[:len(pathWithoutTrailingSlash)-len(".htm")]
	}
	// Need to find out if the path should have a trailing "/". If redirectPath
	// has been set, use taht so that the path we test does NOT have an HTML file
	// extension. Otherwise, use pathWithoutTrailingSlash.
	shouldHaveTrailingSlash := false
	if redirectPath != "" {
		shouldHaveTrailingSlash = srv.checkPathShouldHaveTrailingSlash(redirectPath)
	} else {
		shouldHaveTrailingSlash = srv.checkPathShouldHaveTrailingSlash(pathWithoutTrailingSlash)
	}
	// If the path doesn't have a trailing "/" but should, redirect accordingly.
	if !hasTrailingSlash && shouldHaveTrailingSlash {
		if redirectPath == "" {
			redirectPath = pathWithoutTrailingSlash + "/"
		} else {
			redirectPath = redirectPath + "/"
		}
	}
	// If the path has a trailing "/" but shouldn't, redirect accordingly.
	// (If redirectPath was set earlier, it doesn't need to be set again here.)
	if hasTrailingSlash && !shouldHaveTrailingSlash && redirectPath == "" {
		redirectPath = pathWithoutTrailingSlash
	}
	// If the path does have a trailing "/" and should have that trailing "/",
	// but the redirect path has already been set earlier, add the trailing "/"
	// back in.
	if hasTrailingSlash && shouldHaveTrailingSlash && redirectPath != "" {
		redirectPath = redirectPath + "/"
	}
	// Redirect if we need to.
	if redirectPath != "" {
		http.Redirect(w, r, redirectPath, http.StatusMovedPermanently)
		return
	}

	// At this point, we don't need to redirect.
	// If the path does not end with a file extension, add ".html".
	if !isFileRegexp.MatchString(r.URL.Path) {
		r.URL.Path = pathWithoutTrailingSlash + ".html"
	}
	// Requests to other files should have the correct file extension and can be
	// served as-is.

	// Serve the requested page.
	srv.baseServer.ServeHTTP(w, r)
}

func (srv *CveselServer) checkPathShouldHaveTrailingSlash(path string) bool {
	// Assume "<path>" is the path without a trailing slash or HTML extension.
	// (But other file extensions are allowed because they won't make this
	// function work wrong.)
	// If "<path>" points to a folder and "<path>.html" points to an HTML file,
	// then the user is requesting a collection page and the path should have a
	// trailing "/".
	// Otherwise, there should not be a trailing "/".

	// Remove leading "/" from the path to avoid breaking the I/O functions below.
	path, _ = strings.CutPrefix(path, "/")
	// Try to open the HTML file.
	htmlFile, err := os.OpenInRoot(srv.docRoot, path + ".html")
	// If something went wrong, then the file either doesn't exist or can't
	// be accessed, so it definitely is not a collection page.
	if err != nil {
		return false
	}
	// Don't forget to close it when we're done.
	defer htmlFile.Close()
	// Try to get the HTML file's information.
	d, err := htmlFile.Stat()
	// TODO: Should return an error instead of a result if err != nil.
	if err != nil {
		return false
	}
	// If the HTML file isn't actually a file (it's something else), then it
	// definitely isn't a collection page.
	if !d.Mode().IsRegular() {
		return false
	}
	// Check if it's actually an HTML file. (It can't be a collection page if
	// it isn't a page at all.)
	first512Bytes := make([]byte, 512)
	htmlFile.Read(first512Bytes)
	htmlFileType := http.DetectContentType(first512Bytes)
	if !strings.Contains(htmlFileType, "text/html") {
		return false
	}
	// Try to open the directory.
	collectionDir, err := os.OpenInRoot(srv.docRoot, path)
	// If something went wrong, then the directory either doesn't exist or can't
	// be accessed, so it definitely is not a collection page.
	if err != nil {
		return false
	}
	// Don't forget to close it when we're done.
	defer collectionDir.Close()
	// Try to get the directory's information.
	d, err = collectionDir.Stat()
	// TODO: Should return an error instead of a result if err != nil.
	if err != nil {
		return false
	}
	// If the directory isn't actually a directory (it's something else), then
	// what the user is requesting isn't a collection page.
	if !d.Mode().IsDir() {
		return false
	}
	// If we get all the way here, then the user did request a collection page.
	return true
}

func main() {
	server := NewCveselServer("../client")
	log.Fatal(http.ListenAndServe(":5000", server))
}
