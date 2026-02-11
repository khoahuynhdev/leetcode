package main

import "testing"

func TestLongestBalancedSubarray(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "example 1: all distinct evens and odds",
			nums:     []int{2, 5, 4, 3},
			expected: 4,
		},
		{
			name:     "example 2: duplicates extending balanced subarray",
			nums:     []int{3, 2, 2, 5, 4},
			expected: 5,
		},
		{
			name:     "example 3: best subarray is in the middle",
			nums:     []int{1, 2, 3, 2},
			expected: 3,
		},
		{
			name:     "edge case: single even element",
			nums:     []int{2},
			expected: 0,
		},
		{
			name:     "edge case: single odd element",
			nums:     []int{1},
			expected: 0,
		},
		{
			name:     "edge case: all same even value",
			nums:     []int{2, 2, 2, 2},
			expected: 0,
		},
		{
			name:     "edge case: all same odd value",
			nums:     []int{3, 3, 3},
			expected: 0,
		},
		{
			name:     "one even one odd",
			nums:     []int{1, 2},
			expected: 2,
		},
		{
			name:     "alternating with repeats",
			nums:     []int{1, 2, 1, 2, 3, 4},
			expected: 6,
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
