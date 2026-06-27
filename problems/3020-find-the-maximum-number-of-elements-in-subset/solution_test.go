package main

import "testing"

func TestSolution(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{"example 1: select {4,2,2} -> [2,4,2]", []int{5, 4, 1, 2, 2}, 3},
		{"example 2: all distinct, single element", []int{1, 3, 2, 4}, 1},
		{"edge case: longer chain [2,4,16,4,2]", []int{2, 2, 4, 4, 16}, 5},
		{"edge case: chain stops short, square missing", []int{2, 2, 4, 4}, 3},
		{"edge case: odd run of ones", []int{1, 1, 1}, 3},
		{"edge case: even count of ones reduced to odd", []int{1, 1}, 1},
		{"edge case: two distinct values, no chain", []int{2, 3}, 1},
		{"edge case: large values near 1e9 no overflow (square absent)", []int{1000000000, 1000000000}, 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := maximumLength(tt.nums)
			if result != tt.expected {
				t.Errorf("maximumLength(%v) = %d, want %d", tt.nums, result, tt.expected)
			}
		})
	}
}
