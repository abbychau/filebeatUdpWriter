# FileBeat UDP Writer for ZeroLog AND Gin

[![Go Reference](https://pkg.go.dev/badge/github.com/abbychau/filebeatUdpWriter.svg)](https://pkg.go.dev/github.com/abbychau/filebeatUdpWriter)

FileBeat UDP Logger for ZeroLog and Gin.

# Usage
```go
r := gin.New()

udpLogger, _ := filebeatUdpWriter.CreateLogger("localhost:8125") //you can catch the error here

r.Use(filebeatUdpWriter.GinHandle("gin", udpLogger))
r.Use(filebeatUdpWriter.GinHandle("gin", log.Logger))
```

# Why this Logger?

1. Compatible with ZeroLog and Gin so that you can log in json format for Gin requests.
2. UDP is fast, non-blocking and does not require a file handler.
3. No more docker volume mounting.

# Package Link
https://pkg.go.dev/github.com/abbychau/filebeatUdpWriter