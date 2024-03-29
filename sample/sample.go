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
		fmt.Fprintln(os.Stderr, err)
		return
	}
	resp, err := fetch.New().GetWithContext(context.Background(), u)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	defer resp.Close()
	if _, err := io.Copy(os.Stdout, resp.Body()); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}
