package main

import (
	"flag"
	"fmt"

	"github.com/natemarks/fileshaper/internal"
)

func processArgs() (string, bool) {
	// Define flags
	fileFlag := flag.String("file", "", "File to process")
	sortFlag := flag.Bool("sort", false, "Sort lines by appearance count")
	flag.Parse()
	fmt.Println("sortFlag:", *sortFlag)
	fmt.Println("fileFlag:", *fileFlag)

	return *fileFlag, *sortFlag
}
func main() {
	filename, sort := processArgs()
	lines, err := internal.Lines(filename)
	if err != nil {
		panic(err)
	}

	result := internal.Duplicates(lines)
	if sort {
		fmt.Println("Sorting by appearance count")
		result.SortByLineNumberCount()
	}
	fmt.Println(result.String())
}
