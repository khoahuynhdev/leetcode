package main

import "testing"

func TestNumberOfSpecialChars(t *testing.T) {
	tests := []struct {
		name     string
		word     string
		expected int
	}{
		{"example 1: a, b, c all special", "aaAbcBC", 3},
		{"example 2: no uppercase at all", "abc", 0},
		{"example 3: uppercase precedes lowercase", "AbBCab", 0},
		{"edge case: single lowercase letter", "a", 0},
		{"edge case: single uppercase letter", "Z", 0},
		{"edge case: lowercase after first uppercase", "aAa", 0},
		{"edge case: empty string", "", 0},
		{"edge case: simple special pair", "aA", 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := numberOfSpecialChars(tt.word)
			if result != tt.expected {
				t.Errorf("numberOfSpecialChars(%q) = %d, want %d", tt.word, result, tt.expected)
			}
		})
	}
}
