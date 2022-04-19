package main

import (
	"bufio"
	"os"
	"sort"
)

/*
	Reads lines of text from a file, returning each line as a string in an array of all the lines
*/
func readLinesFromFile(filepath string) ([]string, error) {
	var allLines []string
	file, err := os.Open(filepath)
	if err != nil {
		return allLines, err
	}
	
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		allLines = append(allLines, scanner.Text())
	}
	
	sort.Strings(allLines)
	
	return allLines, nil
}
