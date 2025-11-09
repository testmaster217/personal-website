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
	res, _ := os.ReadFile(fmt.Sprintf("%s%s.html", srv.docRoot, path))
	return string(res)
}

func main() {
	server := &CveselServer{"../client"}
	log.Fatal(http.ListenAndServe(":5000", server))
}