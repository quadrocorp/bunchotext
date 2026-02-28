package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/quadrocorp/bunchotext/internal/core"
)

func main() {
	dirFlag := flag.String("d", ".", "Directory to search")
	outFlag := flag.String("o", "output.txt", "Output file path")
	typeFlag := flag.String("t", "go", fmt.Sprintf("File type preset (%s)", getAvailableTypes()))

	flag.Parse()

	if _, exists := core.FilePatterns[*typeFlag]; !exists {
		fmt.Fprintf(os.Stderr, "error: invalid type '%s'. Available types: %s\n", *typeFlag, getAvailableTypes())
		os.Exit(1)
	}

	fmt.Printf("Scanning directory: %s\n", *dirFlag)
	fmt.Printf("Filtering for type: %s\n", *typeFlag)
	fmt.Printf("Writing to: %s\n", *outFlag)

	if err := core.ProcessDirectory(*dirFlag, *typeFlag, *outFlag); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Done!")
}

func getAvailableTypes() string {
	keys := make([]string, 0, len(core.FilePatterns))
	for k := range core.FilePatterns {
		keys = append(keys, k)
	}
	return strings.Join(keys, ", ")
}
