// +build run

package main

import (
	"fmt"
	"io"
	"os"

	"github.com/spiegel-im-spiegel/fetch"
)

func main() {
	u, err := fetch.URL("https://github.com/spiegel-im-spiegel.gpg")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	resp, err := fetch.New().Get(u)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	defer resp.Body.Close()
	if _, err := io.Copy(os.Stdout, resp.Body); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}
