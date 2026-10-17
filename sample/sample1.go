//go:build run
// +build run

package main

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/goark/fetch"
)

func main() {
	u, err := fetch.URL("https://github.com/spiegel-im-spiegel.gpg")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %+v\n", err)
		return
	}
	resp, err := fetch.New().GetWithContext(
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
	if _, err := io.Copy(os.Stdout, resp.Body()); err != nil {
		fmt.Fprintf(os.Stderr, "error: %+v\n", err)
	}
}
