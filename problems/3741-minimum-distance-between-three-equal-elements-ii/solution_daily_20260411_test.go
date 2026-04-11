package main

import "testing"

func TestMinimumDistance(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "example 1: three 1s close together",
			nums:     []int{1, 2, 1, 1, 3},
			expected: 6,
		},
		{
			name:     "example 2: multiple values with 3+ occurrences",
			nums:     []int{1, 1, 2, 3, 2, 1, 2},
			expected: 8,
		},
		{
			name:     "example 3: single element",
			nums:     []int{1},
			expected: -1,
		},
		{
			name:     "edge case: no value appears three times",
			nums:     []int{1, 2, 1, 2},
			expected: -1,
		},
		{
			name:     "edge case: all elements identical",
			nums:     []int{5, 5, 5, 5, 5},
			expected: 4,
		},
		{
			name:     "edge case: three consecutive identical elements",
			nums:     []int{3, 3, 3},
			expected: 4,
		},
		{
			name:     "edge case: optimal triple not at start of group",
			nums:     []int{1, 2, 1, 2, 2, 2, 1},
			expected: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := minimumDistance(tt.nums)
			if result != tt.expected {
				t.Errorf("minimumDistance(%v) = %d, want %d", tt.nums, result, tt.expected)
			}
		})
	}
}
