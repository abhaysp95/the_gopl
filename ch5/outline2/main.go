package main

import (
  "fmt"
  "os"

  "golang.org/x/net/html"
)

func main() {
  doc, err := html.Parse(os.Stdin)
  if err != nil {
    fmt.Fprintf(os.Stderr, "outline2: %v\n", err)
    os.Exit(1)
  }

  forEachNode(doc, startElement, EndElement)
}

func forEachNode(node *html.Node, pre, post func(n *html.Node)) {
  if pre != nil {  // only function check
    pre(node)
  }

  for c := node.FirstChild; c != nil; c = c.NextSibling {
    forEachNode(c, pre, post)
  }

  if post != nil {
    post(node)
  }
}

var depth int

func startElement(node *html.Node) {
  if node.Type == html.ElementNode {
    fmt.Printf("%*s<%s>\n", depth * 2, "", node.Data)
    depth++
  }
}

func EndElement(node *html.Node) {
  if node.Type == html.ElementNode {
    depth--
    fmt.Printf("%*s</%s>\n", depth * 2, "", node.Data)
  }
}
