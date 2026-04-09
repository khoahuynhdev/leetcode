package main

import "testing"

func TestXorAfterRangeMultiplicationQueries(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		queries  [][]int
		expected int
	}{
		{
			name:     "example 1: all elements multiplied by same value",
			nums:     []int{1, 1, 1},
			queries:  [][]int{{0, 2, 1, 4}},
			expected: 4,
		},
		{
			name:     "example 2: multiple queries with different steps",
			nums:     []int{2, 3, 1, 5, 4},
			queries:  [][]int{{1, 4, 2, 3}, {0, 2, 1, 2}},
			expected: 31,
		},
		{
			name:     "edge case: single element with one query",
			nums:     []int{5},
			queries:  [][]int{{0, 0, 1, 3}},
			expected: 15,
		},
		{
			name:     "edge case: large step skips most elements",
			nums:     []int{1, 2, 3},
			queries:  [][]int{{0, 1, 5, 2}},
			expected: 2 ^ 2 ^ 3, // only index 0 affected: [2, 2, 3], XOR = 3
		},
		{
			name:     "edge case: multiple queries accumulate on same element",
			nums:     []int{1},
			queries:  [][]int{{0, 0, 1, 2}, {0, 0, 1, 3}},
			expected: 6, // 1 * 2 * 3 = 6
		},
		{
			name:     "edge case: multiply by 1 changes nothing",
			nums:     []int{7, 11},
			queries:  [][]int{{0, 0, 1, 1}},
			expected: 7 ^ 11, // [7, 11] unchanged, XOR = 12
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			numsCopy := make([]int, len(tt.nums))
			copy(numsCopy, tt.nums)
			result := xorAfterRangeMultiplicationQueries(numsCopy, tt.queries)
			if result != tt.expected {
				t.Errorf("got %d, want %d", result, tt.expected)
			}
		})
	}
}
