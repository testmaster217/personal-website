package main

import "net/http"

// docRoot is the document root folder for the website.
// Everything that is allowed to be served goes into this folder, and anything
// that's not in this folder should not be served.
const docRoot = "../client/"

func handler(w http.ResponseWriter, r *http.Request) {
	
}

func main() {

}
