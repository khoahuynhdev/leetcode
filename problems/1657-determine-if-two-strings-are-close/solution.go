package solution

import (
	"sort"
)

func closeStrings(word1 string, word2 string) bool {
	// Check if the lengths of both words are equal
	if len(word1) != len(word2) {
		return false
	}

	// Create slices to store the frequency of characters for both words
	freq1 := make([]int, 26)
	freq2 := make([]int, 26)

	// Populate the frequency slices for word1
	for _, char := range word1 {
		freq1[char-'a']++
	}

	// Populate the frequency slices for word2
	for _, char := range word2 {
		freq2[char-'a']++
	}

	// Check if the sets of characters (keys) in both words are the same
	for i := 0; i < 26; i++ {
		if (freq1[i] > 0 && freq2[i] == 0) || (freq1[i] == 0 && freq2[i] > 0) {
			return false
		}
	}

	// Sort the frequency slices in ascending order
	sort.Ints(freq1)
	sort.Ints(freq2)

	// Compare the sorted frequency slices
	for i := 0; i < 26; i++ {
		if freq1[i] != freq2[i] {
			return false
		}
	}

	return true
}
