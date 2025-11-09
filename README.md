# API Gateway

> API Gateway pattern implementation with Gin

Part of the [Modsynth](https://github.com/modsynth) ecosystem.

## Features

- HTTP routing with Gin
- Middleware support
- Request forwarding

## Installation

```bash
go get github.com/modsynth/api-gateway
```

## Quick Start

```go
package main

import "github.com/modsynth/api-gateway"

func main() {
    gw := gateway.New()
    gw.Router().GET("/health", func(c *gin.Context) {
        c.JSON(200, gin.H{"status": "ok"})
    })
    gw.Run(":8080")
}
```

## Version

Current version: `v0.1.0`

## License

MIT
