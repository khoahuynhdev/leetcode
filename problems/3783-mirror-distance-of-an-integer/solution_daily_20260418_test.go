package main

import "testing"

func TestMirrorDistance(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		expected int
	}{
		{"example 1: two-digit number 25", 25, 27},
		{"example 2: trailing zero 10", 10, 9},
		{"example 3: single digit 7", 7, 0},
		{"edge case: smallest input n=1", 1, 0},
		{"edge case: palindrome 121", 121, 0},
		{"edge case: maximum constraint 10^9", 1000000000, 999999999},
		{"edge case: trailing zeros mid-range 1200", 1200, 1179},
		{"edge case: all-nines palindrome 99", 99, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := mirrorDistance(tt.n)
			if got != tt.expected {
				t.Errorf("mirrorDistance(%d) = %d, want %d", tt.n, got, tt.expected)
			}
		})
	}
}
