package main

import "testing"

func TestMinJumps(t *testing.T) {
	tests := []struct {
		name     string
		arr      []int
		expected int
	}{
		{
			name:     "example 1: mixed same-value and adjacent jumps",
			arr:      []int{100, -23, -23, 404, 100, 23, 23, 23, 3, 404},
			expected: 3,
		},
		{
			name:     "example 2: single element, already at last index",
			arr:      []int{7},
			expected: 0,
		},
		{
			name:     "example 3: direct same-value jump from first to last",
			arr:      []int{7, 6, 9, 6, 9, 6, 9, 7},
			expected: 1,
		},
		{
			name:     "edge case: two elements with different values",
			arr:      []int{1, 2},
			expected: 1,
		},
		{
			name:     "edge case: two elements with same value",
			arr:      []int{5, 5},
			expected: 1,
		},
		{
			name:     "edge case: all identical values, single same-value jump suffices",
			arr:      []int{7, 7, 7, 7, 7},
			expected: 1,
		},
		{
			name:     "edge case: strictly distinct values, only +/-1 edges available",
			arr:      []int{1, 2, 3, 4, 5},
			expected: 4,
		},
		{
			name:     "edge case: duplicates that do not shortcut to the end",
			arr:      []int{11, 22, 11, 33, 44},
			expected: 3,
		},
		{
			name:     "edge case: negative values with a useful same-value bucket",
			arr:      []int{-76, 3, 66, -32, 64, 2, -19, -8, -5, -93, 80, -5, -76, -78, 64, 2, 16},
			expected: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := minJumps(tt.arr)
			if got != tt.expected {
				t.Errorf("minJumps(%v) = %d, want %d", tt.arr, got, tt.expected)
			}
		})
	}
}
