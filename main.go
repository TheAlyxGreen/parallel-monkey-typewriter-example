package main

import (
	"sync"
)

const maxThreads = 48
const maxWordLength = 12
const guessesPerThread = 1000000

var foundWordChannel chan string
var letterBookmarks map[string]int
var threadGuard chan struct{}

/*
	This program generates random strings and checks them for a TikTok video I made
*/
func main() {
	
	foundWordChannel = make(chan string)
	letterBookmarks = make(map[string]int)
	
	// the threadGuard is used to cap the number of spawned threads
	threadGuard = make(chan struct{}, maxThreads)
	
	wordList, err := readLinesFromFile("allWords.txt")
	
	if err != nil {
		panic(err)
	}
	
	bookmarkLetterPositions(wordList)
	
	wg := sync.WaitGroup{}
	go foundWordListener(&wg)
	for i := 0; i < maxThreads; i++ {
		wg.Add(1)
		go monkeyGuess(&wg, wordList)
	}
	
	wg.Wait()
}
