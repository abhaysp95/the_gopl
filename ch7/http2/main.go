package main

import (
	"fmt"
	"net/http"
	"net/url"
)

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type database map[string]dollars

func main() {
  db := database{"shoes": 50, "socks": 5}
  http.ListenAndServe("localhost:8000", db)
}
