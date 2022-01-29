package main

import (
	"fmt"
	"log"
	"os"

	"the_gopl/ch5/links"
)

// breadthFirst calls f for each item in worklist.
// any item returned by f are added to worklist.
// f is called atmost once for each item
func breadthFirst(f func(item string) []string, worklist []string) {
  seen := make(map[string]bool)
  for len(worklist) > 0 {
    items := worklist
    worklist = nil
    for _, item := range items {
      if !seen[item] {
        seen[item] = true
        worklist = append(worklist, f(item)...)
      }
    }
  }
}

func crawl(url string) []string {
  fmt.Println(url)
  list, err := links.Extract(url)
  if err != nil {
    log.Print(err)
  }
  return list
}

func main() {
  // Crawl the web breadth-first
  // starting from the command-line arguments
  breadthFirst(crawl, os.Args[1:])
}
