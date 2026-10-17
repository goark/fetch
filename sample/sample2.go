//go:build run
// +build run

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/goark/fetch"
)

func main() {
	u, err := fetch.URL("https://live.staticflickr.com/65535/55582060788_9bf7829404_b.jpg")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %+v\n", err)
		return
	}
	resp, err := fetch.New(fetch.WithHTTP11Client()).GetWithContext(
		context.Background(),
		u,
		fetch.WithRequestHeaderSet("User-Agent", "curl/8.18.0"),
		fetch.WithRequestHeaderSet("Accept", "*/*"),
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %+v\n", err)
		return
	}
	defer resp.Close()
	fmt.Println("HTTP Protocol:", resp.Proto())
}
