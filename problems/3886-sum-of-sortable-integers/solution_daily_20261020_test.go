package main

import "testing"

func TestSumOfSortableIntegers(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{"example 1: single rotation sorts whole array", []int{3, 1, 2}, 3},
		{"example 2: strictly decreasing is never sortable", []int{7, 6, 5}, 0},
		{"example 3: already sorted, every divisor works", []int{5, 8}, 3},
		{"edge case: single element", []int{42}, 1},
		{"edge case: all equal values", []int{4, 4, 4, 4}, 1 + 2 + 4},
		{"edge case: only k=2 works, whole array has too many descents", []int{2, 1, 4, 3}, 2},
		{"edge case: rotated blocks with ordered ranges", []int{2, 1, 4, 3, 6, 5}, 2},
		{"edge case: duplicates at block boundary", []int{2, 1, 2, 3}, 2},
		{"edge case: sorted array of length 6", []int{1, 2, 3, 4, 5, 6}, 1 + 2 + 3 + 6},
		{"edge case: whole-array rotation only", []int{4, 5, 6, 1, 2, 3}, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := sumOfSortableIntegers(tt.nums)
			if result != tt.expected {
				t.Errorf("got %v, want %v", result, tt.expected)
			}
		})
	}
}
