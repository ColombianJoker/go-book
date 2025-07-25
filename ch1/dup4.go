// Dup2 prints the count and text of lines that appear more than once
// in the input. It reads from stdin or from a list of named files,
// and also prints the names of files where duplicated lines occur.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// counts maps a line to its total occurrences
	counts := make(map[string]int)
	// fileOccurrences maps a line to a set (map[string]bool) of filenames where it appeared
	fileOccurrences := make(map[string]map[string]bool)

	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, fileOccurrences, "stdin") // Pass "stdin" as filename for standard input
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, fileOccurrences, arg) // Pass the filename
			f.Close()
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s", n, line)
			if fileset, ok := fileOccurrences[line]; ok {
				fmt.Print("\t(files:")
				for filename := range fileset {
					fmt.Printf(" %s", filename)
				}
				fmt.Println(")")
			} else {
				fmt.Println() // Newline if no files were recorded (e.g., if from stdin only and not handled specifically)
			}
		}
	}
}

// countLines reads lines from a file, updates counts, and records file occurrences.
func countLines(f *os.File, counts map[string]int, fileOccurrences map[string]map[string]bool, filename string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		line := input.Text()
		counts[line]++

		// Initialize the set of files for this line if it doesn't exist
		if fileOccurrences[line] == nil {
			fileOccurrences[line] = make(map[string]bool)
		}
		// Add the current filename to the set for this line
		fileOccurrences[line][filename] = true
	}
	// NOTE: ignoring potential errors from input.Err()
}
