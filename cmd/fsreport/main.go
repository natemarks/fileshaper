package main

import (
	"fmt"
	"github.com/natemarks/fileshaper/internal"
	"os"
)

func getInput() string {
	// Check if any arguments are provided
	if len(os.Args) < 2 {
		panic("Please provide a filename")
	}

	return os.Args[1] // Return the first positional argument
}
func main {
	filename := getInput()
	lines, err := internal.Lines(filename)
	if err != nil {
		panic(err)
	}
	duplicates, err := internal.Duplicates(lines)
	for line, lineNumbers := range duplicates {
		fmt.Printf("%s: %v\n", line, lineNumbers)
	}
}