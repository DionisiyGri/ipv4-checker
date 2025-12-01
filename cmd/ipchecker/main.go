package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/DionisiyGri/ipv4-checker/internal/ipchecker"
)

func main() {
	path := flag.String("file", "", "path to file with IPv4 addresses")
	flag.Parse()

	if *path == "" {
		fmt.Fprintf(os.Stderr, "Usage: ipcounter -file <path/to/file>")
		flag.PrintDefaults()
		os.Exit(2)
	}

	//run the logic to check and count IPs
	res, err := ipchecker.Execute(*path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("lines processed: %d\n", res.Lines)
	fmt.Printf("unique addresses: %d\n", res.Unique)
}
