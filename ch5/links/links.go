package links

import (
  "fmt"
  "net/http"

  "golang.org/x/net/html"
)

// Extract makes the HTTP Get request to the specified URL, parses
// the response as HTML, and returns the link in HTML document
func Extract(url string) ([]string, error) {
  resp, err := http.Get(url)
  if err != nil {
    return nil, err
  }
  if resp.StatusCode != http.StatusOK {
    resp.Body.Close()
    return nil, fmt.Errorf("getting %s: %s", url, resp.Status)
  }

  doc, err := html.Parse(resp.Body)
  resp.Body.Close()
  if err != nil {
    return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
  }

  var links []string
  visitNode := func (node *html.Node) {
    if node.Type == html.ElementNode && node.Data == "a" {
      for _, attr := range node.Attr {
        if attr.Key != "href" {
          continue
        }
        link, err := resp.Request.URL.Parse(attr.Val)
        if err != nil {
          continue  // ignore bad URL
        }
        links = append(links, link.String())
      }
    }
  }
  forEachNode(doc, visitNode, nil)
  return links, nil
}

func forEachNode(node *html.Node, pre, post func(node *html.Node)) {
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
