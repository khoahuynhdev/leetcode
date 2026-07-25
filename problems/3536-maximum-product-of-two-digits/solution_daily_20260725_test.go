package main

import "testing"

func TestSolution(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		expected int
	}{
		{"example 1: n = 31, only one pair 3*1", 31, 3},
		{"example 2: n = 22, same digit reused", 22, 4},
		{"example 3: n = 124, top two digits 2*4", 124, 8},
		{"edge case: smallest input 10, zero is second largest", 10, 0},
		{"edge case: all nines, maximum possible product", 999999999, 81},
		{"edge case: upper bound 10^9, single one and zeros", 1000000000, 0},
		{"edge case: largest digit first, 91", 91, 9},
		{"edge case: largest digit last, 19", 19, 9},
		{"edge case: trailing zero, 90", 90, 0},
		{"edge case: duplicate max separated by zero, 505", 505, 25},
		{"edge case: max digits not adjacent, 918273", 918273, 72},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProduct(tt.n); got != tt.expected {
				t.Errorf("maxProduct(%d) = %v, want %v", tt.n, got, tt.expected)
			}
		})
	}
}
