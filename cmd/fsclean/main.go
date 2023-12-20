package main

import (
	"flag"

	"github.com/natemarks/fileshaper/internal"
)

func processArgs() string {
	// Define flags
	fileFlag := flag.String("file", "", "File to process")
	flag.Parse()

	return *fileFlag
}
func main() {
	filename := processArgs()
	lines, err := internal.Lines(filename)
	if err != nil {
		panic(err)
	}
	result := internal.DeDuplicateLines(lines)
	err = internal.WriteLinesToFile(result, filename+".dedup.txt")
	if err != nil {
		panic(err)
	}
}
