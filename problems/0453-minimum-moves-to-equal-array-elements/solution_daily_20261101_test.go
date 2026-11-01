package main

import "testing"

func TestMinMoves(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{"example 1: increasing sequence", []int{1, 2, 3}, 3},
		{"example 2: all equal", []int{1, 1, 1}, 0},
		{"edge case: single element", []int{42}, 0},
		{"edge case: negative numbers", []int{-1, 0, 1}, 3},
		{"edge case: minimum not first", []int{5, 6, 8, 8, 5}, 7},
		{"edge case: two elements far apart", []int{-1000000000, 1000000000}, 2000000000},
		{"edge case: all negative equal", []int{-7, -7, -7, -7}, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := minMoves(tt.nums)
			if result != tt.expected {
				t.Errorf("minMoves(%v) = %d, want %d", tt.nums, result, tt.expected)
			}
		})
	}
}
