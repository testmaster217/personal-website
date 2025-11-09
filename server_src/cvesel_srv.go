package main

import (
	"fmt"
	"log"
	"net/http"
)

func CveselServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "test page plz ignore")
}

func main() {
	handler := http.HandlerFunc(CveselServer)
	log.Fatal(http.ListenAndServe(":5000", handler))
}