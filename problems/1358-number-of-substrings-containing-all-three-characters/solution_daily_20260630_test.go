package main

import "testing"

func TestNumberOfSubstrings(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		expected int
	}{
		{"example 1: abcabc has 10 valid substrings", "abcabc", 10},
		{"example 2: aaacb has 3 valid substrings", "aaacb", 3},
		{"example 3: minimal abc has 1 valid substring", "abc", 1},
		{"edge case: no b or c means no valid substring", "aaa", 0},
		{"edge case: only two distinct characters", "aabbaa", 0},
		{"edge case: third character only at the end", "aaaaabc", 5},
		{"edge case: long run then all three", "cccccab", 5},
		{"edge case: every character is a valid extension", "cba", 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numberOfSubstrings(tt.s); got != tt.expected {
				t.Errorf("numberOfSubstrings(%q) = %d, want %d", tt.s, got, tt.expected)
			}
		})
	}
}
