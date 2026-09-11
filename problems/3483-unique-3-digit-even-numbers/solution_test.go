package main

import "testing"

func TestSolution(t *testing.T) {
	tests := []struct {
		name     string
		digits   []int
		expected int
	}{
		{"example 1: four distinct digits", []int{1, 2, 3, 4}, 12},
		{"example 2: duplicate digit usable twice", []int{0, 2, 2}, 2},
		{"example 3: all identical digits", []int{6, 6, 6}, 1},
		{"example 4: no even digit available", []int{1, 3, 5}, 0},
		{"edge case: all zeros forces a leading zero", []int{0, 0, 0}, 0},
		{"edge case: single non-zero digit must lead", []int{2, 0, 0}, 1},
		{"edge case: zero allowed in middle and last place", []int{0, 1, 2}, 3},
		{"edge case: every digit even, no duplicates", []int{2, 4, 6, 8}, 24},
		{"edge case: many duplicates collapse to one number", []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 2}, 1},
		{"edge case: repeated digits limit permutations", []int{5, 5, 4}, 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := totalNumbers(tt.digits)
			if result != tt.expected {
				t.Errorf("totalNumbers(%v) = %v, want %v", tt.digits, result, tt.expected)
			}
		})
	}
}
