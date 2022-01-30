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

  ct := resp.Header.Get("Content-Type")
  if ct != "text/html" && !strings.HasPrefix(ct, "text/html;") {
    resp.Body.Close()
    return fmt.Errorf("%s has type %s, not text/html", url, ct)
  }

  doc, err := html.Parse(resp.Body)
  resp.Body.Close()
  if err != nil {
    return fmt.Errorf("parsing %s as HTML: %v", url, err)
  }

  visitNode := func(n *html.Node) {
    if n.Type == html.ElementNode && n.Data == "title" &&
      n.FirstChild != nil {
        fmt.Println(n.FirstChild.Data)
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
    post(nil)
  }
}

func main() {
  for _, arg := range os.Args[1:] {
    if err := title(arg); err != nil {
      log.Fatalf("title: %v", err)
    }
  }
}
