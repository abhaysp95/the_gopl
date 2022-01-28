package github

import (
  "encoding/json"
  "fmt"
  "net/http"
  "net/url"
  "strings"
)

// SearchIssues queries the Github issue tracker
func SearchIssues(terms []string) (*IssueSearchResult, error) {
  q := url.QueryEscape(strings.Join(terms, " "))
  resp, err := http.Get(IssueURL + "?q=" + q)
  if err != nil {
    return nil, err
  }

  // We must close resp.Body on all execution paths
  if resp.StatusCode != http.StatusOK {
    resp.Body.Close()
    return nil, fmt.Errorf("search query failed: %s", resp.Status)
  }

  var result IssueSearchResult
  if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
    resp.Body.Close()
    return nil, err
  }

  resp.Body.Close()
  return &result, nil
}

/**
  https://api.dictionaryapi.dev/api/v2/entries/en/$REPLY
  See if you want to make something for dictionary too
*/
