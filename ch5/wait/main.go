package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
  if len(os.Args) != 2 {
    fmt.Fprintf(os.Stderr, "Usage: ./wait URL\n")
    os.Exit(1)
  }

  url := os.Args[1]
  if err := waitForServer(url); err != nil {
    /* fmt.Fprintf(os.Stderr, "Site is down: %v\n", err)
    os.Exit(1) */
    /* log.SetPrefix("wait: ")
    log.SetFlags(0) */
    log.SetFlags(log.LstdFlags | log.Lshortfile | log.Lmsgprefix)
    log.Fatalf("site is down: %v\n", err)
  }
}

func waitForServer(url string) error {
  const timeout = 1 * time.Minute
  deadline := time.Now().Add(timeout)
  for tries := 0; time.Now().Before(deadline); tries++ {
    _, err := http.Head(url)
    if err == nil {
      return nil  // success
    }
    log.Printf("server not responding (%s); retrying...", err)
    time.Sleep(time.Second << uint(tries))
    // time.Duration is int64 (you can use that too)
  }
  return fmt.Errorf("server %s failed to respond after %s", url, timeout)
}
