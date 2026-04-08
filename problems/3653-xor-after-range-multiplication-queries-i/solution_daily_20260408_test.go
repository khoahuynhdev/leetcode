package main

import "testing"

func TestXorAfterRangeMultiplication(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		queries  [][]int
		expected int
	}{
		{
			name:     "example 1: all ones multiplied by 4",
			nums:     []int{1, 1, 1},
			queries:  [][]int{{0, 2, 1, 4}},
			expected: 4,
		},
		{
			name:     "example 2: two queries with different strides",
			nums:     []int{2, 3, 1, 5, 4},
			queries:  [][]int{{1, 4, 2, 3}, {0, 2, 1, 2}},
			expected: 31,
		},
		{
			name:     "edge case: single element",
			nums:     []int{7},
			queries:  [][]int{{0, 0, 1, 5}},
			expected: 35,
		},
		{
			name:     "edge case: step larger than range",
			nums:     []int{3, 4, 5},
			queries:  [][]int{{0, 2, 3, 2}},
			expected: 6 ^ 4 ^ 5,
		},
		{
			name:     "edge case: no queries",
			nums:     []int{1, 2, 3},
			queries:  [][]int{},
			expected: 1 ^ 2 ^ 3,
		},
		{
			name:     "edge case: large multiplier triggers mod",
			nums:     []int{1000000006},
			queries:  [][]int{{0, 0, 1, 2}},
			expected: (1000000006 * 2) % 1000000007,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// copy nums to avoid mutation across tests
			nums := make([]int, len(tt.nums))
			copy(nums, tt.nums)
			result := xorAfterRangeMultiplication(nums, tt.queries)
			if result != tt.expected {
				t.Errorf("got %v, want %v", result, tt.expected)
			}
		})
	}
}
