package main

import "testing"

func TestSmallestSubsequenceDaily20260719(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"example 1: bcabc", "bcabc", "abc"},
		{"example 2: cbacdcbc", "cbacdcbc", "acdb"},
		{"edge case: single character", "a", "a"},
		{"edge case: all identical characters", "aaaa", "a"},
		{"edge case: already strictly increasing", "abc", "abc"},
		{"edge case: strictly decreasing unique letters", "cba", "cba"},
		{"edge case: leetcode 316 sample bcabc dup", "bbcaac", "bac"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := smallestSubsequence(tt.input)
			if result != tt.expected {
				t.Errorf("smallestSubsequence(%q) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}
