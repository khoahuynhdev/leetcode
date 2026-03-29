package main

import "testing"

func TestCanBeEqual(t *testing.T) {
	tests := []struct {
		name     string
		s1       string
		s2       string
		expected bool
	}{
		{"example 1: swap both pairs", "abcd", "cdab", true},
		{"example 2: impossible transformation", "abcd", "dacb", false},
		{"edge case: already equal", "abcd", "abcd", true},
		{"edge case: all same characters", "aaaa", "aaaa", true},
		{"only even positions need swap", "abcd", "cbad", true},
		{"only odd positions need swap", "abcd", "adcb", true},
		{"even match but odd mismatch", "abcd", "cbda", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := canBeEqual(tt.s1, tt.s2)
			if result != tt.expected {
				t.Errorf("canBeEqual(%q, %q) = %v, want %v", tt.s1, tt.s2, result, tt.expected)
			}
		})
	}
}
