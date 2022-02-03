package main

import (
  "sync"
)

var cache = struct {
  sync.Mutex
  mapping map[string]string
}{
  mapping: make(map[string]string),
}

// safe key lookup
func Lookup(key string) string {
  cache.Lock()
  v := cache.mapping[key]
  cache.Unlock()
  return v
}
