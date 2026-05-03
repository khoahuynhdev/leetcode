package main

import "testing"

func TestRotateString(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		goal     string
		expected bool
	}{
		{"example 1: abcde shifts to cdeab", "abcde", "cdeab", true},
		{"example 2: abcde cannot become abced", "abcde", "abced", false},
		{"edge case: identical strings need zero shifts", "abc", "abc", true},
		{"edge case: different lengths cannot match", "abc", "abcd", false},
		{"edge case: single char equal", "a", "a", true},
		{"edge case: single char different", "a", "b", false},
		{"edge case: same multiset but not a rotation", "aabbcc", "abcabc", false},
		{"edge case: all same characters", "aaaa", "aaaa", true},
		{"edge case: full-length rotation", "abcdefg", "efgabcd", true},
		{"edge case: empty-like minimal mismatch", "ab", "ba", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := rotateString(tt.s, tt.goal)
			if got != tt.expected {
				t.Errorf("rotateString(%q, %q) = %v, want %v", tt.s, tt.goal, got, tt.expected)
			}
		})
	}
}
