package main

import (
	"log"
	"net/http"
	"regexp"
	"strings"
)

var isFileRegexp = regexp.MustCompile(`\.\w+/?$`)

type CveselServer struct {
	// The base file server that my custom server will extend.
	// Should not be set manually, will be set by the NewCveselServer function.
	baseServer http.Handler
}

func NewCveselServer(docRoot string) *CveselServer {
	fileServer := http.FileServer(http.Dir(docRoot))
	return &CveselServer{fileServer}
}

// Should do different things depending on what is requested:
// - If the path ends with a "." followed by anything other than "html" or "htm",
// - - THIS IS CUT: If the file is executable and is not a JS file, it should be executed, and the results should be served.
// - - Otherwise, the file should be served as is.
// - TODO: If the requested path is "/", or "/index[.htm|.html]" (possibly with a "/" at the end), redirect to "/" if necessary, and serve the contents of "index.html".
// - TODO: If this does not apply, and the path ends with ".html" or ".htm" (not counting a trailing "/"), redirect to the equivalent path without the extension.
// - If the path does not end with a file extension (not counting a trailing "/"), serve the file that has that name with ".html" added onto the end.
// - TODO: "Collection" pages (pages whose purpose is to give links to other pages, i.e. "collinvesel.me/blog/"; these will have the same name as a folder in the same directory) should have a trailing "/" in their paths, but other pages should not. Redirect as necessary.
// Whatever gets served after the above are evaluated should be served differently depending on whether the request came from HTMX or not, as indicated by the "HX-Request" header:
// - If the file is NOT an HTML file, serve it as is.
// - PARTIAL TODO; DOESN'T WORK WITH PARTIAL PAGES YET: If the header is msising or set to "false", combine the data to be served with a base page and serve that.
// - TODO: Otherwise, serve the data as is.
//
// TODO: Find out what else I need to make sure this server does, what other standards it needs to comply with, what headers the responses need to have, what the values of those headers should be, etc.
func (srv *CveselServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// If the request path ends with ".html", redirect to the same path without it.
	if strings.HasSuffix(r.URL.Path, ".html") {
		http.Redirect(w, r, r.URL.Path[:len(r.URL.Path) - len(".html")], http.StatusMovedPermanently)
		return
	}
	// Ditto for trailing ".htm"
	if strings.HasSuffix(r.URL.Path, ".htm") {
		http.Redirect(w, r, r.URL.Path[:len(r.URL.Path) - len(".htm")], http.StatusMovedPermanently)
		return
	}

	// If the path does not end with a file extension, add ".html".
	if !isFileRegexp.MatchString(r.URL.Path) {
		newPath, _ := strings.CutSuffix(r.URL.Path, "/")
		r.URL.Path = newPath + ".html"
	}

	srv.baseServer.ServeHTTP(w, r)
}

func main() {
	server := NewCveselServer("../client")
	log.Fatal(http.ListenAndServe(":5000", server))
}
