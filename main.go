package main

import (
	"fmt"
	"os"
)

func main() {
	// if len(os.Args) < 2 {
	// 	fmt.Fprintln(os.Stderr, "no file provided")
	// 	os.Exit(2)
	// }

	// path := os.Args[1]

	f, err := os.Open("ips.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to open %q: %v", "path", err)
		os.Exit(1)
	}
	defer f.Close()
}
