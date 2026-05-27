package main

// Count the Number of Special Characters II
//
// A letter is "special" when it appears in both lowercase and uppercase, and
// every lowercase occurrence comes before the first uppercase occurrence.
//
// In a single pass we record, for each of the 26 letters, the index of the last
// lowercase occurrence and the index of the first uppercase occurrence. A letter
// is special exactly when both exist and lastLower < firstUpper (if the latest
// lowercase precedes the first uppercase, then all lowercase ones do).
//
// Time: O(n), Space: O(1).
func numberOfSpecialChars(word string) int {
	const none = -1

	lastLower := [26]int{}
	firstUpper := [26]int{}
	for i := 0; i < 26; i++ {
		lastLower[i] = none
		firstUpper[i] = none
	}

	for i := 0; i < len(word); i++ {
		ch := word[i]
		switch {
		case ch >= 'a' && ch <= 'z':
			lastLower[ch-'a'] = i
		case ch >= 'A' && ch <= 'Z':
			c := ch - 'A'
			if firstUpper[c] == none {
				firstUpper[c] = i
			}
		}
	}

	count := 0
	for c := 0; c < 26; c++ {
		if lastLower[c] != none && firstUpper[c] != none && lastLower[c] < firstUpper[c] {
			count++
		}
	}
	return count
}
