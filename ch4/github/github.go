// package github provides a Go API for the Github issue tracker
// https://docs.github.com/en/rest/reference/search#search-issues-and-pull-requests

package github

import "time"

const IssueURL  = "https://api.github.com/search/issues"

type IssueSearchResult struct {
  TotalCount int `json:"total_count"`
  Items []*Issue
}

type Issue struct {
  Number int
  HTMLURL string `json:"html_url"`
  Title string
  State string
  User *User
  CreatedAt time.Time `json:"created_at"`
  Body string  // in Markdown format
}

type User struct {
  Login string
  HTMLURL string `json:"html_url"`
}
