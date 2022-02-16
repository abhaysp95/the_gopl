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
      w.WriteHeader(http.StatusNotFound)
      fmt.Fprintf(w, "no such item: %q\n", item)
      return
    }
    fmt.Fprintf(w, "%s: %s\n", item, price)
  default:
    w.WriteHeader(http.StatusNotFound)  // 404
    fmt.Fprintf(w, "no such page: %s\n", req.URL)
  }
}

func main() {
  db := database{"shoes": 50, "socks": 5}
  http.ListenAndServe("localhost:8000", db)
}
