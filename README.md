# [fetch] -- Fetch Data from URL

[![check vulns](https://github.com/goark/fetch/workflows/vulns/badge.svg)](https://github.com/goark/fetch/actions)
[![lint status](https://github.com/goark/fetch/workflows/lint/badge.svg)](https://github.com/goark/fetch/actions)
[![GitHub license](https://img.shields.io/badge/license-Apache%202-blue.svg)](https://raw.githubusercontent.com/goark/fetch/master/LICENSE)
[![GitHub release](https://img.shields.io/github/release/goark/fetch.svg)](https://github.com/goark/fetch/releases/latest)

This package is required Go 1.16 or later.

**Migrated repository to [github.com/goark/fetch][fetch]**

## Import

```go
import "github.com/goark/fetch"
```

## Usage

```go
package main

import (
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
    resp, err := fetch.New().Get(u)
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

[fetch]: https://github.com/goark/fetch "goark/fetch: Fetch Data from URL"
