package main

import (
	"fmt"
	"net/http"
	"net/url"
)

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type database map[string]dollars

func (db database) ServeHTTP(w http.ResponseWriter, req *http.Request) {
  switch req.URL.Path {
  case "/list":
    for item, value := range db {
      fmt.Fprintf(w, "%s: %s\n", item, value)
    }
  case "/price":
    item := req.URL.Query().Get("item")
    // fmt.Println(url.ParseQuery(req.URL.Query().Encode()))
    price, ok := db[item]
    if !ok {
      w.WriteHeader(http.StatusNotFound)  // should be done before writing any message
      fmt.Fprintf(w, "no such item: %q\n", item)
      return
    }
    fmt.Fprintf(w, "%s: %s\n", item, price)
  default:
    /* w.WriteHeader(http.StatusNotFound)  // 404
    fmt.Fprintf(w, "no such page: %s\n", req.URL) */

    // http.ResponseWriter is another interface. It augemnts io.Writer with
    // methods for sending HTTP response headers. Equivantely, we could use:
    msg := fmt.Sprintf("no such page: %s\n", req.URL)
    http.Error(w, msg, http.StatusNotFound)
  }
}

func main() {
  db := database{"shoes": 50, "socks": 5}
  http.ListenAndServe("localhost:8000", db)
}
