package main

import "testing"

func TestSumCounts(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{"example 1: repeated value at both ends", []int{1, 2, 1}, 15},
		{"example 2: all elements equal", []int{1, 1}, 3},
		{"edge case: single element", []int{5}, 1},
		{"edge case: all distinct", []int{1, 2, 3}, 20},
		{"edge case: trailing duplicate", []int{1, 2, 2}, 12},
		{"edge case: max values all equal", []int{100, 100, 100, 100}, 10},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := sumCounts(tt.nums)
			if result != tt.expected {
				t.Errorf("sumCounts(%v) = %d, want %d", tt.nums, result, tt.expected)
			}
		})
	}
}
