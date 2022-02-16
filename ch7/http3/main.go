package main

import (
	"fmt"
	"log"
	"net/http"
)

type dollar float32

func (d dollar) String() string { return fmt.Sprintf("$%.2f", d) }

type database map[string]dollar

// to be used via http.HandlerFunc (adapter to allow the use of ordinary
// functions as HTTP Handler)
func (db database) list(w http.ResponseWriter, req *http.Request) {
  for item, value := range db {
    fmt.Fprintf(w, "%s: %s\n", item, value)
  }
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
  item := req.URL.Query().Get("item")
  price, ok := db[item]
  if !ok {
    w.WriteHeader(http.StatusNotFound)
    fmt.Fprintf(w, "no such item: %q\n", item)
    return
  }
  fmt.Fprintf(w, "%s\n", price)
}

func main() {
  db := database{"shoes": 50, "socks": 5}
  mux := http.NewServeMux()
  mux.Handle("/list", http.HandlerFunc(db.list))  // type conversion
  mux.Handle("/price", http.HandlerFunc(db.price))
  log.Fatal(http.ListenAndServe("localhost:8000", mux))
}
