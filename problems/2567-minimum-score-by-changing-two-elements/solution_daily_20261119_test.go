package main

import "testing"

func TestMinimizeSum(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{"example 1: mixed values", []int{1, 4, 7, 8, 5}, 3},
		{"example 2: three elements", []int{1, 4, 3}, 0},
		{"edge case: all elements equal", []int{5, 5, 5, 5}, 0},
		{"edge case: minimum length with large values", []int{1, 1000000000, 500}, 0},
		{"edge case: outliers on both ends", []int{1, 50, 51, 52, 100}, 2},
		{"edge case: two large outliers", []int{1, 2, 3, 1000, 2000}, 2},
		{"edge case: two small outliers", []int{1, 2, 100, 101, 102}, 2},
		{"edge case: four elements", []int{1, 10, 20, 30}, 9},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := minimizeSum(tt.nums)
			if result != tt.expected {
				t.Errorf("minimizeSum(%v) = %d, want %d", tt.nums, result, tt.expected)
			}
		})
	}
}
