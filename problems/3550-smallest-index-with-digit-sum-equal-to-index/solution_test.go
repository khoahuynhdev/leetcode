package main

import "testing"

func TestSmallestIndex(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{"example 1: match at last index", []int{1, 3, 2}, 2},
		{"example 2: multiple matches, return smallest", []int{1, 10, 11}, 1},
		{"example 3: no match", []int{1, 2, 3}, -1},
		{"edge case: zero at index 0", []int{0}, 0},
		{"edge case: single non-zero element", []int{5}, -1},
		{"edge case: large value 1000 has digit sum 1", []int{7, 1000}, 1},
		{"edge case: zero at index 0 beats later matches", []int{0, 1, 2}, 0},
		{"edge case: match only at index 27", func() []int {
			nums := make([]int, 28)
			for i := range nums {
				nums[i] = 1000
			}
			nums[1] = 5
			nums[27] = 999
			return nums
		}(), 27},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := smallestIndex(tt.nums)
			if result != tt.expected {
				t.Errorf("got %v, want %v", result, tt.expected)
			}
		})
	}
}
