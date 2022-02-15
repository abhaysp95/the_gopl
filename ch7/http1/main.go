package main

import (
	"fmt"
	"log"
	"net/http"
)

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type database map[string]dollars

func main() {
  db := database{"shoes": 50, "socks": 5}
  log.Fatal(http.ListenAndServe("localhost:8000", db))
}

// db will become Handler (interface) when it'll have ServeHTTP method implemented
func (db database) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  for item, price := range db {
    fmt.Fprintf(w, "%s: %s\n", item, price)
  }
}
