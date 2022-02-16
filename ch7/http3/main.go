package main

import (
	"fmt"
	"log"
	"net/http"
)

type dollar float32

func (d dollar) String() string { return fmt.Sprintf("$%.2f", d) }

type database map[string]dollar

func main() {
  db := database{"shoes": 50, "socks": 5}
  mux := http.NewServeMux()
  mux.HandleFunc("/list", http.HandlerFunc(db.list))
  mux.HandleFunc("/price", http.HandlerFunc(db.price))
  log.Fatal(http.ListenAndServe("localhost:8000", mux))
}
