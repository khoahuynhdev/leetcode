package main

import "testing"

func TestSolution(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		expected bool
	}{
		{"example 1: all a's then all b's", "aaabbb", true},
		{"example 2: interleaved a's and b's", "abab", false},
		{"example 3: only b's", "bbb", true},
		{"edge case: single a", "a", true},
		{"edge case: single b", "b", true},
		{"edge case: only a's", "aaaa", true},
		{"edge case: b before a", "ba", false},
		{"edge case: violation at the end", "aabba", false},
		{"edge case: violation at the start", "baaa", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := checkString(tt.s)
			if result != tt.expected {
				t.Errorf("checkString(%q) = %v, want %v", tt.s, result, tt.expected)
			}
		})
	}
}
