package main

import "testing"

func TestMinMoves2(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{"example 1: small consecutive values", []int{1, 2, 3}, 2},
		{"example 2: even length unsorted", []int{1, 10, 2, 9}, 16},
		{"edge case: single element", []int{5}, 0},
		{"edge case: all elements equal", []int{7, 7, 7, 7}, 0},
		{"edge case: two elements", []int{1, 100}, 99},
		{"edge case: negative values", []int{-3, -1, -2}, 2},
		{"edge case: extreme values", []int{-1000000000, 1000000000}, 2000000000},
		{"edge case: outlier does not shift median", []int{1, 0, 0, 8, 6}, 14},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nums := append([]int(nil), tt.nums...)
			result := minMoves2(nums)
			if result != tt.expected {
				t.Errorf("minMoves2(%v) = %d, want %d", tt.nums, result, tt.expected)
			}
		})
	}
}
