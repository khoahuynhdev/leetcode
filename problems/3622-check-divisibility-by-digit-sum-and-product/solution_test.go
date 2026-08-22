package main

import "testing"

func TestSolution(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		expected bool
	}{
		{"example 1: 99 = digit sum 18 + digit product 81", 99, true},
		{"example 2: 23 vs divisor 11 leaves remainder 1", 23, false},
		{"edge case: smallest input n=1, divisor 2", 1, false},
		{"edge case: single digit 9, divisor 18", 9, false},
		{"edge case: zero digit zeroes the product, 10 % 1 == 0", 10, true},
		{"edge case: largest input 1000000, divisor 1", 1000000, true},
		{"edge case: multiple zero digits, 110 % 2 == 0", 110, true},
		{"edge case: three digits divisible, 126 % 21 == 0", 126, true},
		{"edge case: three digits not divisible, 135 vs divisor 24", 135, false},
		{"edge case: all nines 999999 vs divisor 531495", 999999, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := checkDivisibility(tt.n)
			if result != tt.expected {
				t.Errorf("checkDivisibility(%d) = %v, want %v", tt.n, result, tt.expected)
			}
		})
	}
}
