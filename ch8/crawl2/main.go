package main

import (
	"fmt"
	"log"
	"os"

	"the_gopl/ch5/links"
)

// tokens is a counting semaphore used to enforce 20 concurrent request
var tokens = make(chan struct{}, 20)

func crawl(url string) []string {
	fmt.Println(url)
	tokens <- struct{}{}  // acquire a token
	list, err := links.Extract(url)
	<-tokens  // release a token
	if err != nil {
		log.Println(err)
	}
	return list
}

// another problem which crawl1 had was that program doesn't terminate even
// when it has discovered all the links reachable from the initial URL
func main() {
	worklist := make(chan []string)

	var n int  // number of pending sends to worklist

	// start with CLI arg
	n++
	go func() {
		worklist <- os.Args[1:]
	}()

	// crawl the web concurrently
	seen := make(map[string]bool)
	for ; n > 0; n-- {
		list := <-worklist
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				n++
				go func(link string) {
					worklist <- crawl(link)
				}(link)
			}
		}
	}
}
