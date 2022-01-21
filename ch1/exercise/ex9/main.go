// Modify fetch to also print the HTTP status code, found in resp.Status

package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	const prefix = "https://"
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, prefix) {
			url = prefix + url
		}
		resp, err := http.Get(url)
		fmt.Printf("status code: %s = %s\n", url, resp.Status)
		if err != nil {
			fmt.Fprintf(os.Stderr, "ex9: %v\n", err)
			os.Exit(1)
		}
		_, err = io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "ex9: Copy %s: %v\n", url, err)
			os.Exit(1)
		}
	}
}
