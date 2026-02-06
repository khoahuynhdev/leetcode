package main

import "testing"

func TestMinimumRemovals(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		k        int
		expected int
	}{
		{
			name:     "example 1: remove largest element",
			nums:     []int{2, 1, 5},
			k:        2,
			expected: 1,
		},
		{
			name:     "example 2: remove min and max elements",
			nums:     []int{1, 6, 2, 9},
			k:        3,
			expected: 2,
		},
		{
			name:     "example 3: already balanced",
			nums:     []int{4, 6},
			k:        2,
			expected: 0,
		},
		{
			name:     "edge case: single element",
			nums:     []int{5},
			k:        1,
			expected: 0,
		},
		{
			name:     "edge case: all elements same",
			nums:     []int{3, 3, 3, 3},
			k:        1,
			expected: 0,
		},
		{
			name:     "edge case: large k value makes everything balanced",
			nums:     []int{1, 2, 3, 4, 5},
			k:        10,
			expected: 0,
		},
		{
			name:     "edge case: k=1 means only equal elements can coexist",
			nums:     []int{1, 2, 3, 4},
			k:        1,
			expected: 3,
		},
		{
			name:     "edge case: need to remove all but one",
			nums:     []int{1, 100},
			k:        2,
			expected: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := minimumRemovals(tt.nums, tt.k)
			if result != tt.expected {
				t.Errorf("minimumRemovals(%v, %d) = %d, want %d", tt.nums, tt.k, result, tt.expected)
			}
		})
	}
}