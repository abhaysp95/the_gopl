package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func title(url string) error {
  resp, err := http.Get(url)
  if err != nil {
    return err
  }
  defer resp.Body.Close()

  ct := resp.Header.Get("Content-Type")
  if ct != "text/html" && !strings.HasPrefix(ct, "text/html") {
    return fmt.Errorf("%s has %s Content-Type, not text/html", url, ct)
  }

  doc, err := html.Parse(resp.Body)
  if err != nil {
    return fmt.Errorf("parsing %s as HTML: %v", url, err)
  }

  visitNode := func(node *html.Node) {
    if node.Type == html.ElementNode && node.Data == "title" &&
      node.FirstChild != nil {
        fmt.Println(node.FirstChild.Data)
    }
  }

  forEachNode(doc, visitNode, nil)
  return nil
}

func forEachNode(node *html.Node, pre, post func(*html.Node)) {
  if pre != nil {
    pre(node)
  }
  for c := node.FirstChild; c != nil; c = c.NextSibling {
    forEachNode(c, pre, post)
  }
  if post != nil {
    post(node)
  }
}

func main() {
  for _, arg := range os.Args[1:] {
    if err := title(arg); err != nil {
      log.Fatalf("title: %v", err)
    }
  }
}
