package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/natemarks/fileshaper/internal"
)

func processArgs() string {
	// Get the first positional argument
	if flag.NArg() == 0 {
		fmt.Println("Usage: <your_program> [file_path]")
		flag.Usage()
		os.Exit(1)
	}

	return flag.Arg(0)
}
func main() {
	filename := processArgs()
	lines, err := internal.Lines(filename)
	if err != nil {
		panic(err)
	}
	result := internal.DeDuplicateLines(lines)
	internal.WriteLinesToFile(result, filename+".dedup.txt")
}
