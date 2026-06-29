package main

// Problem 1967: Number of Strings That Appear as Substrings in Word
//
// Approach: For each pattern, check whether it appears as a contiguous
// substring of word using strings.Contains, and count the successes.
// The constraints are tiny (<= 100 patterns, lengths <= 100), so a direct
// scan is optimal in practice.

import "strings"

func numOfStrings(patterns []string, word string) int {
	count := 0
	for _, p := range patterns {
		if strings.Contains(word, p) {
			count++
		}
	}
	return count
}
