package main

/*
	This function goes through all the words on the word list and figures out where words for each
	letter of the alphabet begin. This makes checking words faster because, instead of iterating
	over the entire list, it can iterate over just words that begin with a certain letter.
*/
func bookmarkLetterPositions(wordList []string) {
	for i := 0; i < 26; i++ {
		letterBookmarks[string(i+97)] = len(wordList)
	}
	firstLetter := ""
	for lineNumber, word := range wordList {
		if len(word) > 0 {
			firstLetter = word[0:1]
			if letterBookmarks[firstLetter] > lineNumber {
				letterBookmarks[firstLetter] = lineNumber
			}
		}
	}
}
