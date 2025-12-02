# Ip-checker
Count unque IPv4 addresses in a file using bitmap.

## Features
- Streaming reader 
- Fixed 512MB memory bitamp to cover entire IPv4 space
- Tests for core components
- No 3rd party packages/libraries used
- Use custom functions  for trimming and converting to reduce memory allocations

## Installation and run
```go build -o ipchecker ./cmd/ipchecker/main.go``` - build executable binary in project root

```./ipchecker -file /path/to/file``` - run the binary to count unique IPs in the file

```go test ./...``` - execute all tests in all folders


### Test runs
System:
| Component                     | Specification                |
|-------------------------------|-----------------------------|
| Model Name                     | MacBook Pro                 |
| Model Identifier               | MacBookPro13,3              |
| Processor Name                 | Quad-Core Intel Core i7     |
| Processor Speed                | 2.7 GHz                     |
| Number of Processors           | 1                           |
| Total Number of Cores          | 4                           |
| Memory                         | 16 GB                       |

```
/ipcheck -file ip_addresses.txt 
2025/12/02 11:09:12 cant convert ip [194.14] to uint: %!w(*errors.errorString=&{invalid char})
2025/12/02 11:11:03 Lines processed: 7946083138
2025/12/02 11:11:03 Unique addresses: 1000000000
2025/12/02 11:11:03 Memory stats: Alloc=2001.19MB, TotalAlloc=123248.69MB, Sys=2548.36MB, NumGC=240
Processing completed in 14m33.864060655s
```