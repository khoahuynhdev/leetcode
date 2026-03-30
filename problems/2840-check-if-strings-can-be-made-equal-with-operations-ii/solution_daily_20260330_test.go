package main

import "testing"

func TestCheckStrings(t *testing.T) {
	tests := []struct {
		name     string
		s1       string
		s2       string
		expected bool
	}{
		{"example 1: swappable to equal", "abcdba", "cabdab", true},
		{"example 2: not possible", "abe", "bea", false},
		{"edge case: single character equal", "a", "a", true},
		{"edge case: single character different", "a", "b", false},
		{"edge case: already equal strings", "abcdef", "abcdef", true},
		{"edge case: even indices differ", "abab", "bbaa", false},
		{"edge case: all same characters", "aaaa", "aaaa", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := checkStrings(tt.s1, tt.s2)
			if result != tt.expected {
				t.Errorf("checkStrings(%q, %q) = %v, want %v", tt.s1, tt.s2, result, tt.expected)
			}
		})
	}
}
