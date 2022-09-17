package anagram

import (
	"strings"
)

func FindAnagrams(dictionary []string, word string) []string {

	// if empty word return empty slice
	if strings.ReplaceAll(word, " ", "") == "" {
		return []string{}
	}

	res := make([]string, 0, 4)
	// normalize the word
	word = strings.ToLower(strings.ReplaceAll(word, " ", ""))

	// check every word for anagram
	for _, w := range dictionary {
		if isAnagram(w, word) {
			res = append(res, w)
		}
	}
	return res
}

func isAnagram(w, word string) bool {
	// normalize the word
	w = strings.ToLower(strings.ReplaceAll(w, " ", ""))

	// if empty or not of equal length return false
	if w == "" || len(w) != len(word) || w == word {
		return false
	}

	// loop over each rune in one word, if not in the other return false
	// after each iteration delete used rune from second word
	for _, r := range word {
		if strings.ContainsRune(w, r) {
			strings.Replace(w, string(r), "", 1)
			continue
		}
		return false
	}
	return true
}
