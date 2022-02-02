package main

import "fmt"

type Values map[string][]string

// Get returns the first value associated with the given key
// or returns "" if there are none
func (v Values) Get(key string) string {
  if vs := v[key]; len(vs) > 0 {
    return vs[0]
  }
  return ""
}

// Add adds the the value to the key
// It adds to any existing value associated with the key
func (v Values) Add(key, value string) {
  v[key] = append(v[key], value)
}

func main() {
  m := Values{"lang": {"en"}}
  m.Add("item", "1")
  m.Add("item", "2")

  fmt.Println(m.Get("lang"))
  fmt.Println(m.Get("q"))
  fmt.Println(m.Get("item"))
  fmt.Println(m["item"])

  m = nil
  fmt.Println(m.Get("item"))
  m.Add("item", "3")
}
