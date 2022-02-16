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
    http.Error(w, fmt.Sprintf("item not found: %q\n", item), http.StatusNotFound)
    return
  }
  fmt.Fprintf(w, "%s\n", price)
}

func main() {
  db := database{"shoes": 50, "socks": 5}
  // register handler function for the given pattern in DefaultServeMux
  http.HandleFunc("/list", db.list)
  http.HandleFunc("/price", db.price)  // there's also http.Handle() to pass the Handler as 2nd arg
  log.Fatal(http.ListenAndServe("localhost:8000", nil))  // nil means to use DefaultServeMux (global)
}
