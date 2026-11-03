package main

import "testing"

func TestSolution(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{"example 1: mixed values", []int{2, 1, 3}, 3},
		{"example 2: two need raising", []int{4, 4, 5}, 2},
		{"edge case: single element", []int{7}, 0},
		{"edge case: all equal", []int{5, 5, 5, 5}, 0},
		{"edge case: max repeated", []int{3, 1, 3, 2}, 3},
		{"edge case: large spread", []int{100, 1, 1, 1}, 297},
		{"edge case: descending", []int{5, 4, 3, 2, 1}, 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := minMoves(tt.nums)
			if result != tt.expected {
				t.Errorf("got %v, want %v", result, tt.expected)
			}
		})
	}
}
