package main

import (
  "fmt"
  "log"
  "net/http"
)

type dollar float32

func (d dollar) String() string { return fmt.Sprintf("$%.2f", d) }

type database map[string]dollar

func (db database) list(w http.ResponseWriter, req *http.Request) {
  for item, value := range db {
    fmt.Fprintf(w, "%s: %s\n", item, value)
  }
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
  item := req.URL.Query().Get("item")
  price, ok := db[item]
  if !ok {  // item not found
    msg := fmt.Sprintf("no such item: %q\n", item)
    http.Error(w, msg, http.StatusNotFound)
    return
  }
  fmt.Fprintf(w, "%s\n", price)
}

func main() {
  db := database{"shoes": 50, "socks": 5}
  mux := http.NewServeMux()
  mux.HandleFunc("/list", db.list)  // no need for conversion (needed with mux.Handle)
  mux.HandleFunc("/price", db.price)
  log.Fatal(http.ListenAndServe("localhost:8000", mux))
}
