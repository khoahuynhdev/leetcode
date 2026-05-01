package main

import "testing"

func TestMaxRotateFunction(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{"example 1: nums=[4,3,2,6] best at k=3", []int{4, 3, 2, 6}, 26},
		{"example 2: single element returns 0", []int{100}, 0},
		{"edge case: all zeros", []int{0, 0, 0, 0}, 0},
		{"edge case: all negative values", []int{-2, -1, -3}, -4},
		{"edge case: mixed signs picks middle rotation", []int{1, -1, 2, -2}, 5},
		{"edge case: two elements", []int{1, 2}, 2},
		{"edge case: optimum not at k=0 or k=n-1", []int{-2, 3, 1, -5, 4}, 13},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := append([]int(nil), tt.nums...)
			got := maxRotateFunction(input)
			if got != tt.expected {
				t.Errorf("maxRotateFunction(%v) = %d, want %d", tt.nums, got, tt.expected)
			}
		})
	}
}
