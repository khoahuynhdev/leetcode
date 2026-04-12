package main

import "testing"

func TestMinimumDistance(t *testing.T) {
	tests := []struct {
		name     string
		word     string
		expected int
	}{
		{"example 1: CAKE", "CAKE", 3},
		{"example 2: HAPPY", "HAPPY", 6},
		{"edge case: two distinct letters can use both fingers for free", "AZ", 0},
		{"edge case: same letter repeated costs zero", "AAAA", 0},
		{"edge case: three-letter word with two free starts", "ABC", 1},
		{"edge case: same-row sequence with both free starts", "ABCDEF", 4},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := minimumDistance(tt.word)
			if got != tt.expected {
				t.Errorf("minimumDistance(%q) = %d, want %d", tt.word, got, tt.expected)
			}
		})
	}
}
