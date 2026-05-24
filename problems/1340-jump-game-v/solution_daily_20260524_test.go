package main

import "testing"

func TestSolution(t *testing.T) {
	tests := []struct {
		name     string
		arr      []int
		d        int
		expected int
	}{
		{
			name:     "example 1: jump 10->8->6->7 visits 4 indices",
			arr:      []int{6, 4, 14, 6, 8, 13, 9, 7, 10, 6, 12},
			d:        2,
			expected: 4,
		},
		{
			name:     "example 2: all equal, no jumps possible",
			arr:      []int{3, 3, 3, 3, 3},
			d:        3,
			expected: 1,
		},
		{
			name:     "example 3: strictly decreasing, visit everything",
			arr:      []int{7, 6, 5, 4, 3, 2, 1},
			d:        1,
			expected: 7,
		},
		{
			name:     "edge case: single element cannot jump",
			arr:      []int{5},
			d:        1,
			expected: 1,
		},
		{
			name:     "edge case: two elements, jump from taller to shorter",
			arr:      []int{1, 2},
			d:        1,
			expected: 2,
		},
		{
			name:     "edge case: strictly increasing chains via left jumps",
			arr:      []int{1, 2, 3, 4},
			d:        2,
			expected: 4,
		},
		{
			name:     "edge case: peak in middle reaches both sides",
			arr:      []int{1, 3, 1},
			d:        1,
			expected: 2,
		},
		{
			name:     "edge case: tall bar blocks a closer shorter bar within d",
			arr:      []int{4, 6, 1},
			d:        2,
			expected: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := maxJumps(tt.arr, tt.d)
			if result != tt.expected {
				t.Errorf("maxJumps(%v, %d) = %d, want %d", tt.arr, tt.d, result, tt.expected)
			}
		})
	}
}
