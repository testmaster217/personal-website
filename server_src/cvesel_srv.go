package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

type CveselServer struct {
	docRoot string
}

func (srv *CveselServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, srv.getPage(r.URL.Path))
}

func (srv *CveselServer) getPage(path string) string {
	// Should do different things depending on what is requested:
	// - If the path ends with a "." followed by anything other than "html" or "htm", it should serve the file as is.
	// - If the path ends with ".html", ".htm", or "/", redirect to the equivalent path without the trailing thing.
	// - If the path ends with anything else, serve the file that has that name with ".html" added onto the end.
	// TODO: Add special treatment for the homepage/"index.html".
	// TODO: Add treatment for HTMX vs non-HTMX requests.
	res, _ := os.ReadFile(fmt.Sprintf("%s%s.html", srv.docRoot, path))
	return string(res)
}

func main() {
	server := &CveselServer{"../client"}
	log.Fatal(http.ListenAndServe(":5000", server))
}
