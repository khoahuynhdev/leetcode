package main

import "testing"

func TestRotatedDigitsDaily20260502(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		expected int
	}{
		{"example 1: n=10 yields {2,5,6,9}", 10, 4},
		{"example 2: n=1 has no good numbers", 1, 0},
		{"example 3: n=2 yields just {2}", 2, 1},
		{"edge case: n=8 still {2,5,6} since 8 maps to itself", 8, 3},
		{"edge case: n=9 yields {2,5,6,9}", 9, 4},
		{"edge case: n=20 includes two-digit goods", 20, 9},
		{"edge case: n=30 skips 23-24, 27, 30", 30, 15},
		{"edge case: upper constraint n=10000", 10000, 2320},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := rotatedDigitsDaily20260502(tt.n)
			if got != tt.expected {
				t.Errorf("rotatedDigitsDaily20260502(%d) = %d, want %d", tt.n, got, tt.expected)
			}
		})
	}
}

func TestIsGoodDaily20260502(t *testing.T) {
	tests := []struct {
		name     string
		x        int
		expected bool
	}{
		{"1 is not good (self-rotates)", 1, false},
		{"2 is good (rotates to 5)", 2, true},
		{"3 is invalid", 3, false},
		{"8 is not good (self-rotates)", 8, false},
		{"10 is not good (only self-rotating digits)", 10, false},
		{"23 is invalid due to 3", 23, false},
		{"25 is good", 25, true},
		{"101 is not good (all self-rotating)", 101, false},
		{"269 is good", 269, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isGoodDaily20260502(tt.x)
			if got != tt.expected {
				t.Errorf("isGoodDaily20260502(%d) = %v, want %v", tt.x, got, tt.expected)
			}
		})
	}
}
