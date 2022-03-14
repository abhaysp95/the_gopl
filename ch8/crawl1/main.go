package main

import (
	"fmt"
	"log"
	"os"
	"the_gopl/ch5/links"
)

func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Println(err)
	}
	return list
}

func main() {
	worklist := make(chan []string)

	// start with command-line arguments
	// will block the main goroutine (that's why it got its own personal
	// goroutine)
	go func() {
		worklist <- os.Args[1:]
	}()

	// crawl the web concurrently
	seen := make(map[string]bool)
	for list := range worklist {
		fmt.Println("list: ", list)
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				go func(link string) {
					worklist <- crawl(link)
				}(link)
			}
		}
	}
}
