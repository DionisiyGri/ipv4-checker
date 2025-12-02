package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"
	"time"

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

	startTime := time.Now()

	//run the logic to check and count IPs
	res, err := ipchecker.Execute(*path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	log.Printf("Lines processed: %d\n", res.Lines)
	log.Printf("Unique addresses: %d\n", res.Unique)

	printMemStats()
	fmt.Printf("Processing completed in %v", time.Since(startTime))
}

func printMemStats() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	log.Printf("Memory stats: Alloc=%.2fMB, TotalAlloc=%.2fMB, Sys=%.2fMB, NumGC=%d",
		float64(m.Alloc)/1024/1024,
		float64(m.TotalAlloc)/1024/1024,
		float64(m.Sys)/1024/1024,
		m.NumGC,
	)
}
