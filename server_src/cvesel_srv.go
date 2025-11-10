package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"
)

var isFileRegexp = regexp.MustCompile(`\.\w+/?$`)
// var isHtml = regexp.MustCompile(`\.html?/?$`)

type CveselServer struct {
	docRoot string
}

func (srv *CveselServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, srv.getPage(r.URL.Path))
}

func (srv *CveselServer) getPage(path string) string {
	// Should do different things depending on what is requested:
	// - If the path ends with a "." followed by anything other than "html" or "htm", 
	// - - If the file is executable and is not a JS file, it should be executed, and the results should be served.
	// - - Otherwise, the file should be served as is.
	// - If the requested path is empty or "/index[.htm|.html]" (possibly with a "/" at the end), redirect to an empty path if necessary, and serve the contents of "index.html".
	// - If this does not apply, and the path ends with ".html", ".htm", or "/", redirect to the equivalent path without the trailing thing.
	// - If the path ends with anything else, serve the file that has that name with ".html" added onto the end.
	// Whatever gets served after the above are evaluated should be served differently depending on whether the request came from HTMX or not, as indicated by the "HX-Request" header:
	// - If the header is msising or set to "false", combine the data to be served with a base page and serve that.
	// - Otherwise, serve the data as is.
	if isFileRegexp.MatchString(path) {
		res, _ := os.ReadFile(fmt.Sprintf("%s%s", srv.docRoot, path))
		return string(res)
	} else {
		res, _ := os.ReadFile(fmt.Sprintf("%s%s.html", srv.docRoot, path))
		return string(res)
	}
}

func main() {
	server := &CveselServer{"../client"}
	log.Fatal(http.ListenAndServe(":5000", server))
}
