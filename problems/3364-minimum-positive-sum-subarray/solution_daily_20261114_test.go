package main

import "testing"

func TestSolution(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		l, r     int
		expected int
	}{
		{"example 1: mixed signs, best is [3,-2]", []int{3, -2, 1, 4}, 2, 3, 1},
		{"example 2: no positive-sum subarray", []int{-2, 2, -3, 1}, 2, 3, -1},
		{"example 3: all positive, shortest window wins", []int{1, 2, 3, 4}, 2, 4, 3},
		{"edge case: single positive element", []int{5}, 1, 1, 5},
		{"edge case: single negative element", []int{-5}, 1, 1, -1},
		{"edge case: zero sums are not positive", []int{0, 0, 0}, 1, 3, -1},
		{"edge case: only full-length window allowed", []int{4, -1, -2}, 3, 3, 1},
		{"edge case: longer window gives smaller positive sum", []int{5, -4, 3}, 1, 3, 1},
		{"edge case: all negative", []int{-1, -2, -3}, 1, 3, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := minimumSumSubarray(tt.nums, tt.l, tt.r)
			if got != tt.expected {
				t.Errorf("minimumSumSubarray(%v, %d, %d) = %d, want %d", tt.nums, tt.l, tt.r, got, tt.expected)
			}
		})
	}
}
