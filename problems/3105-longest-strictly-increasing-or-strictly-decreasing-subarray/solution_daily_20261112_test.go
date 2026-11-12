package main

import "testing"

func TestSolution(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{"example 1: mixed with equal neighbors", []int{1, 4, 3, 3, 2}, 2},
		{"example 2: all equal", []int{3, 3, 3, 3}, 1},
		{"example 3: fully decreasing", []int{3, 2, 1}, 3},
		{"edge case: single element", []int{7}, 1},
		{"edge case: fully increasing", []int{1, 2, 3, 4, 5}, 5},
		{"edge case: zigzag", []int{1, 3, 2, 4, 3}, 2},
		{"edge case: longest run at end", []int{5, 5, 1, 2, 3, 4}, 4},
		{"edge case: decreasing then increasing", []int{5, 4, 3, 2, 1, 2, 3}, 5},
		{"edge case: plateau splits run", []int{1, 2, 2, 3, 4}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := longestMonotonicSubarray(tt.nums)
			if got != tt.expected {
				t.Errorf("longestMonotonicSubarray(%v) = %d, want %d", tt.nums, got, tt.expected)
			}
		})
	}
}
