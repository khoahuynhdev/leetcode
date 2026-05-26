package main

// numberOfSpecialChars counts letters that appear in both lowercase and
// uppercase form in word. We make a single pass and record, for each of the
// 26 letters, whether it was seen in lowercase and whether it was seen in
// uppercase. A letter is special when both flags are set, so the answer is the
// size of the intersection of the two sets.
func numberOfSpecialChars(word string) int {
	var lower, upper [26]bool

	for _, c := range word {
		switch {
		case c >= 'a' && c <= 'z':
			lower[c-'a'] = true
		case c >= 'A' && c <= 'Z':
			upper[c-'A'] = true
		}
	}

	count := 0
	for i := 0; i < 26; i++ {
		if lower[i] && upper[i] {
			count++
		}
	}
	return count
}
