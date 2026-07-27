package main

import "testing"

func TestMaxProduct(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{"example 1: distinct values, max pair is 5 and 4", []int{3, 4, 5, 2}, 12},
		{"example 2: duplicate maximum used twice", []int{1, 5, 4, 5}, 16},
		{"example 3: minimum length array", []int{3, 7}, 12},
		{"edge case: all elements identical", []int{2, 2, 2, 2}, 1},
		{"edge case: minimum possible values yield zero", []int{1, 1}, 0},
		{"edge case: maximum constraint values", []int{1000, 1000}, 998001},
		{"edge case: maximum appears first, exercises demotion", []int{9, 1, 1}, 0},
		{"edge case: strictly descending input", []int{5, 4, 3, 2, 1}, 12},
		{"edge case: strictly ascending input", []int{1, 2, 3, 4, 5}, 12},
		{"edge case: one large value among ones", []int{1, 1, 1000, 1}, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := maxProduct(tt.nums)
			if result != tt.expected {
				t.Errorf("maxProduct(%v) = %v, want %v", tt.nums, result, tt.expected)
			}
		})
	}
}
