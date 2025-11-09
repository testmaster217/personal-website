package main

import (
	"fmt"
	"log"
	"net/http"
)

func CveselServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, GetPage(r.URL.Path))
}

func GetPage(path string) string {
	if path == "/testpage" {
		return "test page plz ignore"
	}

	if path == "/testpage2" {
		return "2nd test page plz ignore"
	}

	return ""
}

func main() {
	handler := http.HandlerFunc(CveselServer)
	log.Fatal(http.ListenAndServe(":5000", handler))
}