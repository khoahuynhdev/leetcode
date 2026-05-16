package main

import "testing"

func TestFindMinDaily20260516(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{"example 1: [1,3,5] no duplicates", []int{1, 3, 5}, 1},
		{"example 2: [2,2,2,0,1] duplicates before pivot", []int{2, 2, 2, 0, 1}, 0},
		{"edge case: single element", []int{1}, 1},
		{"edge case: all elements equal", []int{3, 3, 3, 3, 3}, 3},
		{"edge case: not rotated (ascending)", []int{1, 2, 3, 4, 5}, 1},
		{"edge case: duplicates straddling pivot", []int{10, 1, 10, 10, 10}, 1},
		{"edge case: duplicates on both sides of pivot", []int{1, 1, 1, 0, 1}, 0},
		{"edge case: two elements rotated", []int{2, 1}, 1},
		{"edge case: two elements not rotated", []int{1, 2}, 1},
		{"edge case: negative numbers rotated", []int{-1, -7, -5, -3}, -7},
		{"edge case: rotated once", []int{5, 1, 2, 3, 4}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findMin(tt.nums)
			if got != tt.expected {
				t.Errorf("findMin(%v) = %d, want %d", tt.nums, got, tt.expected)
			}
		})
	}
}
