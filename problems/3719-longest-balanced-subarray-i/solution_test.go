package main

import "testing"

func TestLongestBalancedSubarray(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "example 1: array with 2 distinct evens and 2 distinct odds",
			nums:     []int{2, 5, 4, 3},
			expected: 4,
		},
		{
			name:     "example 2: array with duplicates",
			nums:     []int{3, 2, 2, 5, 4},
			expected: 5,
		},
		{
			name:     "example 3: smaller balanced subarray",
			nums:     []int{1, 2, 3, 2},
			expected: 3,
		},
		{
			name:     "edge case: single element",
			nums:     []int{5},
			expected: 0,
		},
		{
			name:     "edge case: all even numbers",
			nums:     []int{2, 4, 6, 8},
			expected: 0,
		},
		{
			name:     "edge case: all odd numbers",
			nums:     []int{1, 3, 5, 7},
			expected: 0,
		},
		{
			name:     "edge case: two elements one even one odd",
			nums:     []int{2, 3},
			expected: 2,
		},
		{
			name:     "edge case: alternating with duplicates",
			nums:     []int{1, 1, 2, 2},
			expected: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := longestBalancedSubarray(tt.nums)
			if result != tt.expected {
				t.Errorf("longestBalancedSubarray(%v) = %d, want %d", tt.nums, result, tt.expected)
			}
		})
	}
}
