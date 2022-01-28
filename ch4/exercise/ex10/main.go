// catagories issues based on date (less than a month old, less than a year old, more than a year old)

package main

import (
  "fmt"
  "flag"
  "log"
  "os"
  "time"

  "the_gopl/ch4/github"
)

var cBef = flag.Int("b", 0, "Issue before number of months")

func main() {
  flag.Parse()

  res, err := github.SearchIssues(os.Args[1:])
  if err != nil {
    log.Fatal(err)
  }

  if *cBef == 0 {
    printItem(res.Items)
  } else {
    reqTime := time.Now().AddDate(0, -(*cBef), 0)
    selectedIssues := make([]*github.Issue, 0)

    for _, item := range res.Items {
      if item.CreatedAt.After(reqTime) {
        selectedIssues = append(selectedIssues, item)
      }
    }
    printItem(selectedIssues)
  }
}

func printItem(items []*github.Issue) {
  format := "#%-5d %9.9s %.55s"

  for _, item := range items {
    fmt.Println(format, item.Number, item.User.Login, item.Title)
  }
}
