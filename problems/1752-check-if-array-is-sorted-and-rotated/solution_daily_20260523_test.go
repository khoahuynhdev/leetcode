package main

import "testing"

func TestSolution(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected bool
	}{
		{"example 1: rotated sorted array", []int{3, 4, 5, 1, 2}, true},
		{"example 2: not a rotation of any sorted array", []int{2, 1, 3, 4}, false},
		{"example 3: already sorted, rotation by zero", []int{1, 2, 3}, true},
		{"edge case: single element", []int{1}, true},
		{"edge case: all duplicates", []int{1, 1, 1}, true},
		{"edge case: strictly decreasing", []int{3, 2, 1}, false},
		{"edge case: two equal elements", []int{2, 2}, true},
		{"edge case: rotation with duplicates", []int{6, 10, 6}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := check(tt.nums)
			if result != tt.expected {
				t.Errorf("check(%v) = %v, want %v", tt.nums, result, tt.expected)
			}
		})
	}
}
