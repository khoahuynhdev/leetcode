package main

import (
	"strings"
	"testing"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		name     string
		num      int
		sum      int
		expected string
	}{
		{"example 1: small sum fits in one digit", 2, 3, "30"},
		{"example 2: nines first then remainder", 2, 17, "98"},
		{"example 3: sum too large for digit count", 1, 10, ""},
		{"edge case: single digit, sum one", 1, 1, "1"},
		{"edge case: all nines exactly", 3, 27, "999"},
		{"edge case: one more than max", 3, 28, ""},
		{"edge case: remainder in middle followed by zeros", 5, 20, "99200"},
		{"edge case: sum one with many digits", 4, 1, "1000"},
		{"edge case: large input all nines", 200000, 1800000, strings.Repeat("9", 200000)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := maxSumOfSquares(tt.num, tt.sum)
			if result != tt.expected {
				if len(result) > 50 || len(tt.expected) > 50 {
					t.Errorf("maxSumOfSquares(%d, %d) mismatch (len got %d, want %d)", tt.num, tt.sum, len(result), len(tt.expected))
				} else {
					t.Errorf("maxSumOfSquares(%d, %d) = %q, want %q", tt.num, tt.sum, result, tt.expected)
				}
			}
		})
	}
}
