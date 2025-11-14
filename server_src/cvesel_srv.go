package main

import (
	"log"
	"net/http"
	"strings"
	"regexp"
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
// - TODO: If the requested path is empty, "/", or "/index[.htm|.html]" (possibly with a "/" at the end), redirect to an empty path if necessary, and serve the contents of "index.html".
// - TODO: If this does not apply, and the path ends with ".html", ".htm", or "/", redirect to the equivalent path without the trailing thing.
// - If the path does not end with a file extension or a "/", serve the file that has that name with ".html" added onto the end.
// Whatever gets served after the above are evaluated should be served differently depending on whether the request came from HTMX or not, as indicated by the "HX-Request" header:
// - If the file is NOT an HTML file, serve it as is.
// - PARTIAL TODO; DOESN'T WORK WITH PARTIAL PAGES YET: If the header is msising or set to "false", combine the data to be served with a base page and serve that.
// - TODO: Otherwise, serve the data as is.
//
// TODO: Make sure that all files served have the correct MIME type. (Apparently, the CSS for the real pages wasn't being applied properly bc the MIME type was wrong, but the page itself and the images were fine. ¯\_(ツ)_/¯)
// TODO: Find out what else I need to make sure this server does, what other standards it needs to comply with, what headers the responses need to have, what the values of those headers should be, etc.
func (srv *CveselServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// If the path does not end with a file extension, change it to one that ends with ".html".
	if !isFileRegexp.MatchString(r.URL.Path) {
		newPath, _ := strings.CutSuffix(r.URL.Path, "/")
		newPath = strings.Join([]string{newPath, ".html"}, "")
		r.URL.Path = newPath
	}

	srv.baseServer.ServeHTTP(w, r)
}

func main() {
	server := NewCveselServer("../client")
	log.Fatal(http.ListenAndServe(":5000", server))
}
