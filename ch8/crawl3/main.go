package main

import (
	"fmt"
	"log"
	"os"

	"the_gopl/ch5/links"
)

func crawl(url string) []string {
	fmt.Println(url)
	links, err := links.Extract(url)
	if err != nil {
		log.Println(err)
	}
	return links
}

func main() {
	worklist := make(chan []string)  // list of URLs, may have duplicate
	unseenLinks := make(chan string)  // de-duplicate URLs

	// Add command-line arguments to worklist
	go func() {
		worklist <- os.Args[1:]
	}()

	// create 20 crawler goroutines to fetch each unseen link
	for i := 0; i < 20; i++ {
		go func() {
			for link := range unseenLinks {
				foundLinks := crawl(link)
				go func() {
					worklist <- foundLinks
				}()
			}
		}()
	}

	// The main goroutine de-duplicates worklist items and sends the unseen
	// ones to the crawler
	seen := make(map[string]bool)
	for list := range worklist {
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				unseenLinks <- link
			}
		}
	}
}
