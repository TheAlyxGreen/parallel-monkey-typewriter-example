package main

import (
	"fmt"
	"sync"
)

/*
	Listens to the found word channel to receive new words from the guesser threads, which
	it then counts and prints out.
*/
func foundWordListener(wg *sync.WaitGroup) {
	currentWord := <-foundWordChannel
	uniqueWordsList := make(map[string]int)
	
	uniqueWordCount := 0
	totalWordCount := 0
	
	for currentWord != "" {
		totalWordCount++
		uniqueWordsList[currentWord] = 1            // only checked to see if it exists, so value doesn't matter
		if len(uniqueWordsList) > uniqueWordCount { // check if a new word was added to the list
			uniqueWordCount = len(uniqueWordsList)
			fmt.Printf(
				"Found unique word %d: %s\t(%d total words found)\n",
				uniqueWordCount,
				currentWord,
				totalWordCount,
			)
		}
		currentWord = <-foundWordChannel // get next word
	}
	wg.Done()
}
