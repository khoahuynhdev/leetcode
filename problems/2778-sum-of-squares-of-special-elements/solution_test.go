package main

import "testing"

func TestSumOfSquares(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{"example 1: n=4, special positions 1,2,4", []int{1, 2, 3, 4}, 21},
		{"example 2: n=6, special positions 1,2,3,6", []int{2, 7, 1, 19, 18, 3}, 63},
		{"edge case: single element", []int{7}, 49},
		{"edge case: prime length only first and last", []int{3, 50, 50, 50, 4}, 25},
		{"edge case: two elements both special", []int{5, 6}, 61},
		{"edge case: perfect square length", []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 1 + 9 + 81},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := sumOfSquares(tt.nums)
			if result != tt.expected {
				t.Errorf("sumOfSquares(%v) = %d, want %d", tt.nums, result, tt.expected)
			}
		})
	}
}
