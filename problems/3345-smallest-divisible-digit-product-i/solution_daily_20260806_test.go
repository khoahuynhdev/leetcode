package main

import "testing"

func TestSmallestNumber(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		t        int
		expected int
	}{
		{"example 1: n contains a zero digit so its product 0 is divisible by 2", 10, 2, 10},
		{"example 2: 15 has product 5, next candidate 16 has product 6", 15, 3, 16},
		{"edge case: t = 1 divides every product, answer is n itself", 1, 1, 1},
		{"edge case: t = 10 forces a scan up to a number containing 0", 1, 10, 10},
		{"edge case: rollover across the 99 -> 100 boundary", 99, 7, 100},
		{"edge case: upper bound n = 100 already has digit product 0", 100, 9, 100},
		{"edge case: single digit n is its own answer when it divides t", 5, 5, 5},
		{"edge case: single digit scan 7 -> 8 for t = 4", 7, 4, 8},
		{"edge case: 19 has product 9, rolls over to 20 with product 0", 19, 7, 20},
		{"edge case: smallest possible input n = 1, t = 2", 1, 2, 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := smallestNumber(tt.n, tt.t); got != tt.expected {
				t.Errorf("smallestNumber(%d, %d) = %v, want %v", tt.n, tt.t, got, tt.expected)
			}
		})
	}
}

func TestDigitProduct(t *testing.T) {
	tests := []struct {
		name     string
		x        int
		expected int
	}{
		{"single digit", 7, 7},
		{"contains a zero digit", 10, 0},
		{"two nonzero digits", 16, 6},
		{"three digits with zeros", 100, 0},
		{"all nines", 99, 81},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := digitProduct(tt.x); got != tt.expected {
				t.Errorf("digitProduct(%d) = %v, want %v", tt.x, got, tt.expected)
			}
		})
	}
}
