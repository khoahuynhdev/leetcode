package main

import "testing"

func TestSolution(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		k        int
		expected int
	}{
		{
			name:     "example 1: first stable index is the last one",
			nums:     []int{5, 0, 1, 4},
			k:        3,
			expected: 3,
		},
		{
			name:     "example 2: no index is stable",
			nums:     []int{3, 2, 1},
			k:        1,
			expected: -1,
		},
		{
			name:     "example 3: single element always scores zero",
			nums:     []int{0},
			k:        0,
			expected: 0,
		},
		{
			name:     "edge case: all elements equal, score is zero everywhere",
			nums:     []int{7, 7, 7},
			k:        0,
			expected: 0,
		},
		{
			name:     "edge case: strictly increasing array is stable at index 0",
			nums:     []int{1, 2, 3, 4},
			k:        0,
			expected: 0,
		},
		{
			name:     "edge case: score dips then rises, so no early stop or binary search",
			nums:     []int{5, 0, 2, 9, 1},
			k:        4,
			expected: 2,
		},
		{
			name:     "edge case: only the final index qualifies when k is zero",
			nums:     []int{1, 0, 2},
			k:        0,
			expected: 2,
		},
		{
			name:     "edge case: two decreasing elements with k too small",
			nums:     []int{2, 1},
			k:        0,
			expected: -1,
		},
		{
			name:     "edge case: maximum constraint values, k exactly on the boundary",
			nums:     []int{1000000000, 0},
			k:        1000000000,
			expected: 0,
		},
		{
			name:     "edge case: maximum constraint values, k one below the boundary",
			nums:     []int{1000000000, 0},
			k:        999999999,
			expected: -1,
		},
		{
			name:     "edge case: single large element with k zero",
			nums:     []int{1000000000},
			k:        0,
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := smallestStableIndex(tt.nums, tt.k); got != tt.expected {
				t.Errorf("smallestStableIndex(%v, %d) = %d, want %d", tt.nums, tt.k, got, tt.expected)
			}
		})
	}
}
