package main

import "testing"

func TestNumberOfSpecialChars(t *testing.T) {
	tests := []struct {
		name     string
		word     string
		expected int
	}{
		{"example 1: aaAbcBC has special a, b, c", "aaAbcBC", 3},
		{"example 2: abc has no uppercase", "abc", 0},
		{"example 3: abBCab only b is special", "abBCab", 1},
		{"edge case: single lowercase char", "a", 0},
		{"edge case: single uppercase char", "Z", 0},
		{"edge case: all uppercase only", "ABC", 0},
		{"edge case: same letter both cases", "Aa", 1},
		{"edge case: repeated chars do not inflate count", "aaaAAA", 1},
		{"edge case: full alphabet in both cases", "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ", 26},
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
