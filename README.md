# [fetch] -- Fetch Data from URL

[![check vulns](https://github.com/spiegel-im-spiegel/fetch/workflows/vulns/badge.svg)](https://github.com/spiegel-im-spiegel/fetch/actions)
[![lint status](https://github.com/spiegel-im-spiegel/fetch/workflows/lint/badge.svg)](https://github.com/spiegel-im-spiegel/fetch/actions)
[![GitHub license](https://img.shields.io/badge/license-Apache%202-blue.svg)](https://raw.githubusercontent.com/spiegel-im-spiegel/fetch/master/LICENSE)
[![GitHub release](https://img.shields.io/github/release/spiegel-im-spiegel/fetch.svg)](https://github.com/spiegel-im-spiegel/fetch/releases/latest)

## Import

```go
import "github.com/spiegel-im-spiegel/fetch"
```

## Usage

```go
package main

import (
    "context"
    "fmt"
    "io"
    "net/http"
    "os"

    "github.com/spiegel-im-spiegel/fetch"
)

func main() {
    u, err := fetch.URL("https://github.com/spiegel-im-spiegel.gpg")
    if err != nil {
        fmt.Fprintln(os.Stderr, err)
        return
    }
    resp, err := fetch.New(
        fetch.WithHTTPClient(&http.Client{}),
    ).Get(
        u,
        fetch.WithContext(context.Background()),
    )
    if err != nil {
        fmt.Fprintln(os.Stderr, err)
        return
    }
    defer resp.Close()
    if _, err := io.Copy(os.Stdout, resp.Body()); err != nil {
        fmt.Fprintln(os.Stderr, err)
    }
}
```

## Modules Requirement Graph

[![dependency.png](./dependency.png)](./dependency.png)

[fetch]: https://github.com/spiegel-im-spiegel/fetch "spiegel-im-spiegel/fetch: Fetch Data from URL"
