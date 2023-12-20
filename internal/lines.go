package internal

import (
	"fmt"
	"os"
	"strings"
)

type DuplicateLineNumbers struct {
	Line        string
	LineNumbers []int
}
type DuplicateSet struct {
	Lines []DuplicateLineNumbers
}

func (d *DuplicateSet) Contains(line string) bool {
	for _, l := range d.Lines {
		if l.Line == line {
			return true
		}
	}
	return false
}

func (d *DuplicateSet) String() string {
	var result []string
	for _, l := range d.Lines {
		result = append(result, l.Line+": "+fmt.Sprint(l.LineNumbers))
	}
	return strings.Join(result, "\n")
}

// Lines returns the lines of the given file.
func Lines(filename string) ([]string, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(content), "\n")
	return lines, nil
}

// DeDuplicateLines removes duplicate lines from the input.
func DeDuplicateLines(input []string) []string {
	uniqueElements := make(map[string]bool)
	var result []string
	if len(input) == 0 {
		return []string{}
	}
	for _, element := range input {
		if !uniqueElements[element] {
			uniqueElements[element] = true
			result = append(result, element)
		}
	}

	return result
}

// Duplicates returns a map of duplicate lines and their line numbers.
func Duplicates(input []string) (result DuplicateSet) {
	if len(input) == 0 {
		return DuplicateSet{Lines: []DuplicateLineNumbers{}}
	}
	// create map of unique lines and their indexes
	elementLines := make(map[string][]int)
	for idx, element := range input {
		if val, ok := elementLines[element]; ok {
			elementLines[element] = append(val, idx)
		} else {
			elementLines[element] = []int{idx}
		}
	}

	// create the result in the order of the input lines
	for _, line := range input {
		if !result.Contains(line) {
			result.Lines = append(result.Lines, DuplicateLineNumbers{Line: line, LineNumbers: elementLines[line]})
		}
	}
	return result
}

// WriteLinesToFile writes the lines to the given file path.
func WriteLinesToFile(lines []string, filePath string) error {
	content := strings.Join(lines, "\n")

	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(content)
	if err != nil {
		return err
	}

	return nil
}
