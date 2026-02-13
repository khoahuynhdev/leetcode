package main

import "testing"

func TestLongestBalancedSubstring(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		expected int
	}{
		// Examples from the problem
		{"example 1: abbac", "abbac", 4},
		{"example 2: aabcc", "aabcc", 3},
		{"example 3: aba", "aba", 2},
		// Edge cases
		{"single character", "a", 1},
		{"all same character", "aaaa", 4},
		{"all three equal", "abc", 3},
		{"all three repeated equally", "aabbcc", 6},
		{"two chars alternating", "abab", 4},
		{"no balanced longer than 1", "abccc", 3},
		{"entire string balanced with three chars", "abcabc", 6},
		{"long single char run", "bbbbb", 5},
		{"pair balanced at end", "cccab", 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := longestBalancedSubstring(tt.s)
			if result != tt.expected {
				t.Errorf("longestBalancedSubstring(%q) = %d, want %d", tt.s, result, tt.expected)
			}
		})
	}
}
