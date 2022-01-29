package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func main() {
  for _, url := range os.Args[1:] {
    links, err := findLinks(url)
    if err != nil {
      fmt.Fprintf(os.Stderr, "findlinks2: %v\n", err)
      os.Exit(1)
    }
    for _, link := range links {
      fmt.Println(link)
    }
  }
}

// findLinks performs HTTP Get on url, parses the html and extracts and returns
// the link
func findLinks(url string) ([]string, error) {
  resp, err := http.Get(url)
  if err != nil {
    return nil, err
  }
  if resp.StatusCode != http.StatusOK {
    resp.Body.Close()
    return nil, fmt.Errorf("getting %s: %s\n", url, resp.Status)
  }
  doc, err := html.Parse(resp.Body)
  resp.Body.Close()
  if err != nil {
    return nil, fmt.Errorf("parsing %s as HTML: %v\n", url, err)
  }
  return visit(nil, doc), nil
}

func visit(links []string, node *html.Node) []string {
  if node.Type == html.ElementNode && node.Data == "a" {
    for _, attr := range node.Attr {
      if attr.Key == "href" {
        links = append(links, attr.Val)
      }
    }
  }
  for c := node.FirstChild; c != nil; c = c.NextSibling {
    links = visit(links, c)
  }
  return links
}
