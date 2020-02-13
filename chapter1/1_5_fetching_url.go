package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	url := strings.Join(os.Args[1:], " ")
	if !strings.HasPrefix(url, "http://") { // Exercise 1.8
		url = "http://" + url
	}
	resp, err := http.Get(url)
	fmt.Println("HTTP status code:", resp.Status) // Exercise 1.9
	if err != nil {
		fmt.Fprintf(os.Stderr, "fech: %v\n", err)
		os.Exit(1)
	}
	b, err := io.Copy(os.Stdout, resp.Body) // Exercise 1.7
	resp.Body.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
		os.Exit(1)
	}
	fmt.Printf("%v", b)
}
