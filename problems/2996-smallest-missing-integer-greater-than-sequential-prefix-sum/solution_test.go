package main

import "testing"

func TestMissingInteger(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{"example 1: prefix sum itself is missing", []int{1, 2, 3, 2, 5}, 6},
		{"example 2: must skip past 12, 13, 14", []int{3, 4, 5, 1, 12, 14, 13}, 15},
		{"edge case: single element", []int{1}, 2},
		{"edge case: answer exceeds every value", []int{50}, 51},
		{"edge case: first pair breaks the run", []int{5, 1, 2, 3}, 6},
		{"edge case: entire array is sequential", []int{1, 2, 3, 4, 5}, 15},
		{"edge case: duplicates after the prefix", []int{1, 2, 1, 2, 3, 4}, 5},
		{"edge case: descending array, prefix is one element", []int{4, 3, 2, 1}, 5},
		{"edge case: repeated first element", []int{7, 7, 7}, 8},
		{"edge case: max sum from full 1..50 sequence", fullSequence(50), 1275},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := missingInteger(tt.nums)
			if got != tt.expected {
				t.Errorf("missingInteger(%v) = %v, want %v", tt.nums, got, tt.expected)
			}
		})
	}
}

// fullSequence returns [1, 2, ..., n].
func fullSequence(n int) []int {
	nums := make([]int, n)
	for i := range nums {
		nums[i] = i + 1
	}
	return nums
}
