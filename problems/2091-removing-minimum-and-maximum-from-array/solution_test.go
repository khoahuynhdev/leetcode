package main

import "testing"

func TestSolution(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "example 1: split, 2 from front and 3 from back",
			nums:     []int{2, 10, 7, 5, 4, 1, 8, 6},
			expected: 5,
		},
		{
			name:     "example 2: min and max adjacent near the front",
			nums:     []int{0, -4, 19, 1, 8, -2, -3, 5},
			expected: 3,
		},
		{
			name:     "example 3: single element is both min and max",
			nums:     []int{101},
			expected: 1,
		},
		{
			name:     "edge case: two elements, must delete both",
			nums:     []int{1, 2},
			expected: 2,
		},
		{
			name:     "edge case: min and max at opposite ends, split wins",
			nums:     []int{1, 3, 4, 2, 5},
			expected: 2,
		},
		{
			name:     "edge case: strictly decreasing, max first and min last",
			nums:     []int{5, 4, 3, 2, 1},
			expected: 2,
		},
		{
			name:     "edge case: min and max both in the middle, front sweep wins",
			nums:     []int{3, 1, 5, 2, 4},
			expected: 3,
		},
		{
			name:     "edge case: all negative values",
			nums:     []int{-5, -1, -3},
			expected: 2,
		},
		{
			name:     "edge case: max before min near the end, back sweep wins",
			nums:     []int{3, 4, 5, 6, 7, 9, 8, 2},
			expected: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := minimumDeletions(tt.nums)
			if result != tt.expected {
				t.Errorf("minimumDeletions(%v) = %v, want %v", tt.nums, result, tt.expected)
			}
		})
	}
}
