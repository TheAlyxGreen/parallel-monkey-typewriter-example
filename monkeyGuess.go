package main

import (
	"math/rand"
	"sync"
)

/*
	This thread handles the guessing by generating a random string, then checking if it's a word
*/
func monkeyGuess(wg *sync.WaitGroup, wordList []string) {
	for guessNumber := 0; guessNumber < guessesPerThread; guessNumber++ {
		guess := getRandomString()
		// spinning this off into a separate thread makes things faster for some reason
		wg.Add(1)
		threadGuard <- struct{}{}
		go validateWord(wg, guess, wordList, letterBookmarks)
	}
	<-threadGuard
	wg.Done()
}

/*
	Generate a random string of letters a-z between 1 and maxWordLength long
*/
func getRandomString() string {
	wordLen := rand.Intn(maxWordLength-1) + 1
	word := ""
	for i := 0; i < wordLen; i++ {
		word = word + string(rand.Intn(25)+97)
	}
	return word
}

/*
	Check a given word against the word list using the bookmarks as a lookup table
*/
func validateWord(
	wg *sync.WaitGroup,
	word string,
	wordlist []string,
	bookmarks map[string]int,
) bool {
	wordFirstLetter := word[0:1]
	lookupStart := bookmarks[wordFirstLetter]
	
	for i := lookupStart; i <= len(wordlist); i++ {
		entry := wordlist[i]
		entryFirstLetter := entry[0:1]
		
		if entry == word {
			foundWordChannel <- word
			<-threadGuard
			wg.Done()
			return true
		} else if entryFirstLetter != wordFirstLetter {
			<-threadGuard
			wg.Done()
			return false
		}
		
	}
	<-threadGuard
	wg.Done()
	return false
}
