package main

import "testing"

func TestGetMinDistance(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		target   int
		start    int
		expected int
	}{
		{
			name:     "example 1: target at far right of start",
			nums:     []int{1, 2, 3, 4, 5},
			target:   5,
			start:    3,
			expected: 1,
		},
		{
			name:     "example 2: single element matches",
			nums:     []int{1},
			target:   1,
			start:    0,
			expected: 0,
		},
		{
			name:     "example 3: all equal, start is the best match",
			nums:     []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			target:   1,
			start:    0,
			expected: 0,
		},
		{
			name:     "edge case: multiple matches, pick nearest to start",
			nums:     []int{1, 5, 3, 5, 7, 5, 9},
			target:   5,
			start:    4,
			expected: 1, // index 3 or 5, both distance 1
		},
		{
			name:     "edge case: match before start",
			nums:     []int{7, 2, 9, 4, 5},
			target:   7,
			start:    4,
			expected: 4,
		},
		{
			name:     "edge case: start index itself is the target",
			nums:     []int{4, 8, 2, 6, 10},
			target:   6,
			start:    3,
			expected: 0,
		},
		{
			name:     "edge case: target at beginning, start at end",
			nums:     []int{3, 1, 2, 4, 5},
			target:   3,
			start:    4,
			expected: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getMinDistance(tt.nums, tt.target, tt.start)
			if got != tt.expected {
				t.Errorf("getMinDistance(%v, %d, %d) = %d; want %d",
					tt.nums, tt.target, tt.start, got, tt.expected)
			}
		})
	}
}
