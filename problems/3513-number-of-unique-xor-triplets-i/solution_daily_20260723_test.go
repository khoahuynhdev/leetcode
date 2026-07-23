package main

import "testing"

func TestSolution(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{"example 1: nums = [1,2] -> {1,2}", []int{1, 2}, 2},
		{"example 2: nums = [3,1,2] -> {0,1,2,3}", []int{3, 1, 2}, 4},
		{"edge case: single element n=1", []int{1}, 1},
		{"edge case: n=4 power of two boundary", []int{4, 3, 2, 1}, 8},
		{"edge case: n=7 fills 3-bit range", []int{7, 6, 5, 4, 3, 2, 1}, 8},
		{"edge case: n=8 crosses to 4-bit range", []int{8, 7, 6, 5, 4, 3, 2, 1}, 16},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := uniqueXorTriplets(tt.nums)
			if result != tt.expected {
				t.Errorf("uniqueXorTriplets(%v) = %d, want %d", tt.nums, result, tt.expected)
			}
		})
	}
}
