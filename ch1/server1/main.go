// server1 is minimal "echo" server

package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)  // each request calls handler
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// handler echos the path component of request URL r
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path: %q\n", r.URL.Path)
}
