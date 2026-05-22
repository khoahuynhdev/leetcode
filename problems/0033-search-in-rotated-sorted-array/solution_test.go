package main

import "testing"

func TestSolution(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		target   int
		expected int
	}{
		{"example 1: target in rotated tail", []int{4, 5, 6, 7, 0, 1, 2}, 0, 4},
		{"example 2: target absent", []int{4, 5, 6, 7, 0, 1, 2}, 3, -1},
		{"example 3: single element, target absent", []int{1}, 0, -1},
		{"edge case: single element, target present", []int{1}, 1, 0},
		{"edge case: empty array", []int{}, 5, -1},
		{"edge case: not rotated, target present", []int{1, 2, 3, 4, 5}, 4, 3},
		{"edge case: target at last index", []int{4, 5, 6, 7, 0, 1, 2}, 2, 6},
		{"edge case: target at first index", []int{4, 5, 6, 7, 0, 1, 2}, 4, 0},
		{"edge case: rotated by one, target at start", []int{5, 1, 2, 3, 4}, 5, 0},
		{"edge case: two elements, target in second", []int{3, 1}, 1, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := search(tt.nums, tt.target)
			if result != tt.expected {
				t.Errorf("search(%v, %d) = %d, want %d", tt.nums, tt.target, result, tt.expected)
			}
		})
	}
}
