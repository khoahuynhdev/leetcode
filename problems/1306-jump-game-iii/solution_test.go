package main

import "testing"

func TestCanReach(t *testing.T) {
	tests := []struct {
		name     string
		arr      []int
		start    int
		expected bool
	}{
		{
			name:     "example 1: reach zero via index 5 -> 4 -> 1 -> 3",
			arr:      []int{4, 2, 3, 0, 3, 1, 2},
			start:    5,
			expected: true,
		},
		{
			name:     "example 2: reach zero from index 0 via 0 -> 4 -> 1 -> 3",
			arr:      []int{4, 2, 3, 0, 3, 1, 2},
			start:    0,
			expected: true,
		},
		{
			name:     "example 3: zero exists but unreachable from start 2",
			arr:      []int{3, 0, 2, 1, 2},
			start:    2,
			expected: false,
		},
		{
			name:     "edge case: start is already on a zero",
			arr:      []int{0, 1, 2},
			start:    0,
			expected: true,
		},
		{
			name:     "edge case: single element must be zero",
			arr:      []int{0},
			start:    0,
			expected: true,
		},
		{
			name:     "edge case: no zero anywhere in the array",
			arr:      []int{1, 1, 1, 1},
			start:    0,
			expected: false,
		},
		{
			name:     "edge case: two elements, jump from index 1 to zero at index 0",
			arr:      []int{0, 1},
			start:    1,
			expected: true,
		},
		{
			name:     "edge case: visited guard prevents infinite loop when no zero is reachable",
			arr:      []int{1, 1},
			start:    0,
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := canReach(tt.arr, tt.start)
			if got != tt.expected {
				t.Errorf("canReach(%v, %d) = %v, want %v", tt.arr, tt.start, got, tt.expected)
			}
		})
	}
}
