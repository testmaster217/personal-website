package main

import (
	"fmt"
	"log"
	"net/http"
)

func CveselServer(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/testpage" {
		fmt.Fprintf(w, "test page plz ignore")
		return
	}

	if r.URL.Path == "/testpage2" {
		fmt.Fprintf(w, "2nd test page plz ignore")
		return
	}
}

func main() {
	handler := http.HandlerFunc(CveselServer)
	log.Fatal(http.ListenAndServe(":5000", handler))
}