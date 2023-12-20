package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/natemarks/fileshaper/internal"
)

func processArgs() (string, bool) {
	// Define flags
	sortFlag := flag.Bool("sort", false, "Sort lines by appearance count")
	flag.Parse()

	// Get the first positional argument
	if flag.NArg() == 0 {
		fmt.Println("Usage: <your_program> [file_path] [-sort]")
		flag.Usage()
		os.Exit(1)
	}

	return flag.Arg(0), *sortFlag
}
func main() {
	filename, sort := processArgs()
	lines, err := internal.Lines(filename)
	if err != nil {
		panic(err)
	}

	result := internal.Duplicates(lines)
	if sort {
		result.SortByLineNumberCount()
	}
	fmt.Println(result)
}
