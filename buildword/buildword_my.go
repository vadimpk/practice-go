package buildword

type fragment struct {
	position int
	length   int
}

func BuildWord(word string, fragments []string) int {
	fs := makeFragments(word, fragments)

	min := 999999
	r(0, len(word), fs, 0, &min)

	return min % 999999
}

// recursive function that finds all possible ways to build a word and defines the minimum one
func r(position, length int, fragments []fragment, count int, min *int) {
	// if the word is built, check if min and reset all vars
	if position == length {
		if count < *min {
			*min = count
		}
		count = 0
		position = 0
	} else {
		// find fragments to be put next
		fs := findFragments(position, length, fragments)
		if len(fs) > 0 {
			// loop through each fragment and call recursion
			for _, f := range fs {
				r(position+f.length, length, fragments, count+1, min)
			}
			// if no fragments found reset the vars
		} else {
			count = 0
			position = 0
		}
	}
}

// finds all possible fragments that can be put into word with given length at given position
// from list of all fragments
func findFragments(position int, length int, fragments []fragment) []fragment {
	res := make([]fragment, 0)
	for _, f := range fragments {
		if f.position == position && position+f.length <= length {
			res = append(res, f)
		}
	}
	return res
}

// create list of fragments from list of strings
func makeFragments(word string, fragments []string) []fragment {
	res := make([]fragment, 0)
	for _, fr := range fragments {
		for _, p := range isSub(word, fr) {
			res = append(res, fragment{position: p, length: len(fr)})
		}
	}
	return res
}

// returns all possible positions of sub in word
// TODO can be optimised
func isSub(word, sub string) []int {

	runes := []rune(sub)
	pos := 0

	res := make([]int, 0)

	for i, r := range word {
		if pos == len(runes) {
			res = append(res, i-pos)
			pos = 0
		}
		if r == runes[pos] {
			pos++
		} else {
			pos = 0
		}
	}
	if pos != 0 && pos == len(runes) {
		res = append(res, len(word)-pos)
	}
	return res
}
