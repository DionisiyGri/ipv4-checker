# Ip-checker
Count unque IPv4 addresses in a file using bitmap.

## Features
- Streaming reader
- Fixed 512MB memory bitamp to cover entire IPv4 space
- Tests for core components
  

## Installation and run
```go build -o ipchecker ./cmd/ipchecker/main.go``` - build executable binary in project root
```./ipchecker -file /path/to/file``` - run the binary to count unique IPs in the file

```go test ./...``` - execute all tests in all folders