package main

import (
	"fmt"

	"golang.org/x/net/html"
)

func soleTitle(doc *html.Node) (title string, err error) {
  type bailout struct{}

  defer func() {
    switch p := recover(); p {
    case nil:
      // no panic
    case bailout{}:
      // "expected" panic
      err = fmt.Errorf("multiple title elements")
    default:
      panic(p)
    }
  }()

  // bail out of recursion if we find more than one non-empty title element
  forEachNode(doc, func(n *html.Node) {
    if n.Type == html.ElementNode && n.Data == "title" && n.FirstChild != nil {
      if title != "" {
        panic(bailout{})
      }
      title = n.FirstChild.Data
    }
  }, nil)

  if title == "" {
    return "", fmt.Errorf("no title element")
  }

  return title, nil
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
