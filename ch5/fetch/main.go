package main

import (
	"io"
	"net/http"
	"os"
	"path"
)

func fetch(url string) (filename string, n int64, err error) {
  resp, err := http.Get(url)
  if err != nil {
    return "", 0, err  // I don't suppose this is going to be any different then just "return"
    // It could be that you want to return something other named variable of
    // error type other than err and so you just can't write "return err", but
    // you have to put all the return parameters, but here it is returning err,
    // so a simple "return" should do the trick
  }
  defer resp.Body.Close()

  local := path.Base(resp.Request.URL.Path)
  if local == "/" {
    local = "index.html"
  }

  f, err := os.Create(local)
  if err != nil {
    return "", 0, err
  }

  n, err = io.Copy(f, resp.Body)
  if err != nil && err != io.EOF {
    return "", 0, err
  }

  if closeErr := f.Close(); err == nil {
    err = closeErr
  }

  return local, n, err
}
