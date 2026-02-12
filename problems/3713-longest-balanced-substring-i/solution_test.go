package main

import "testing"

func TestLongestBalancedSubstring(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		expected int
	}{
		{"example 1: abbac", "abbac", 4},
		{"example 2: zzabccy", "zzabccy", 4},
		{"example 3: aba", "aba", 2},
		{"edge case: single character", "a", 1},
		{"edge case: all same characters", "aaaa", 4},
		{"edge case: all distinct characters", "abcde", 5},
		{"edge case: two characters alternating", "ababab", 6},
		{"edge case: long balanced at end", "xyzaabbcc", 6},
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
