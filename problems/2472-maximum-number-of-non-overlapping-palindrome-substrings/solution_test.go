package main

import (
	"strings"
	"testing"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		k        int
		expected int
	}{
		{"example 1: pick \"aba\" and \"dbbd\"", "abaccdbbd", 3, 2},
		{"example 2: no palindrome of length at least k", "adbcda", 2, 0},
		{"edge case: k is 1 so every character counts", "abc", 1, 3},
		{"edge case: single character with k of 1", "a", 1, 1},
		{"edge case: k larger than the string", "abc", 4, 0},
		{"edge case: empty input", "", 1, 0},
		{"edge case: k equals the string length and it is a palindrome", "aaaaa", 5, 1},
		{"edge case: repeated characters split into pairs", "aaaaa", 2, 2},
		{"edge case: repeated characters split into triples", "aaaaaaaaaa", 3, 3},
		{"edge case: shorter k on example 1 finds \"aba\", \"cc\", \"bb\"", "abaccdbbd", 2, 3},
		{"edge case: splitting \"abba\" into \"bb\" and \"aa\" is no worse", "abbaaa", 2, 2},
		{"edge case: all candidates overlap around one center", "racecar", 3, 1},
		{"edge case: nested palindromes favour the shorter pick", "abacaba", 3, 2},
		{"edge case: only an even-length palindrome qualifies", "leetcode", 2, 1},
		{"edge case: no palindrome of length 3 exists", "abc", 3, 0},
		{"edge case: k of 1 on a palindrome-free string", "adbcda", 1, 6},
		{"edge case: long uniform string at maximum constraint", strings.Repeat("a", 2000), 2, 1000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxPalindromes(tt.s, tt.k); got != tt.expected {
				t.Errorf("maxPalindromes(%q, %d) = %v, want %v", tt.s, tt.k, got, tt.expected)
			}
		})
	}
}
