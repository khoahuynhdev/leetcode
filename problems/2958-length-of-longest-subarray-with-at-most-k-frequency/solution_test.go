package main

import "testing"

func TestMaxSubarrayLength(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		k        int
		expected int
	}{
		{
			name:     "example 1: repeating 1,2,3 pattern with k=2",
			nums:     []int{1, 2, 3, 1, 2, 3, 1, 2},
			k:        2,
			expected: 6,
		},
		{
			name:     "example 2: alternating 1,2 with k=1 allows only distinct pairs",
			nums:     []int{1, 2, 1, 2, 1, 2, 1, 2},
			k:        1,
			expected: 2,
		},
		{
			name:     "example 3: all identical elements capped by k",
			nums:     []int{5, 5, 5, 5, 5, 5, 5},
			k:        4,
			expected: 4,
		},
		{
			name:     "edge case: single element",
			nums:     []int{7},
			k:        1,
			expected: 1,
		},
		{
			name:     "edge case: all distinct values, whole array is good",
			nums:     []int{4, 8, 15, 16, 23, 42},
			k:        1,
			expected: 6,
		},
		{
			name:     "edge case: k larger than any frequency, never shrinks",
			nums:     []int{1, 1, 2, 2, 3, 3},
			k:        6,
			expected: 6,
		},
		{
			name:     "edge case: repeats separated by a gap force a multi-step shrink",
			nums:     []int{1, 2, 3, 4, 1},
			k:        1,
			expected: 4,
		},
		{
			name:     "edge case: large values beyond array indexing range",
			nums:     []int{1000000000, 999999999, 1000000000, 999999999},
			k:        2,
			expected: 4,
		},
		{
			name:     "edge case: best window sits at the end of the array",
			nums:     []int{9, 9, 9, 1, 2, 3, 4},
			k:        1,
			expected: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := maxSubarrayLength(tt.nums, tt.k)
			if result != tt.expected {
				t.Errorf("maxSubarrayLength(%v, %d) = %v, want %v", tt.nums, tt.k, result, tt.expected)
			}
		})
	}
}
