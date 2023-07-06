# FileBeat UDP Writer for ZeroLog AND Gin

[![Go Reference](https://pkg.go.dev/badge/github.com/abbychau/filebeatUdpWriter.svg)](https://pkg.go.dev/github.com/abbychau/filebeatUdpWriter)

FileBeat log writer for GIN-zerolog-udp-filebeat stack.

# Usage
```go
r := gin.New()

udpLogger, _ := filebeatUdpWriter.CreateLogger("localhost:8125") //you can catch the error here

r.Use(filebeatUdpWriter.GinHandle("gin", udpLogger))
r.Use(filebeatUdpWriter.GinHandle("gin", log.Logger))
```

# Package Link
https://pkg.go.dev/github.com/abbychau/filebeatUdpWriter