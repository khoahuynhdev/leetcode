package main

import "testing"

func TestMinJumps(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "example 1: adjacent step then teleport via prime 2",
			nums:     []int{1, 2, 4, 6},
			expected: 2,
		},
		{
			name:     "example 2: adjacent step then teleport via prime 3",
			nums:     []int{2, 3, 4, 7, 9},
			expected: 2,
		},
		{
			name:     "example 3: no useful teleportation, walk all the way",
			nums:     []int{4, 6, 5, 8},
			expected: 3,
		},
		{
			name:     "edge case: single element starts at destination",
			nums:     []int{1},
			expected: 0,
		},
		{
			name:     "edge case: two elements, one adjacent step",
			nums:     []int{1, 1},
			expected: 1,
		},
		{
			name:     "edge case: only ones, must walk every step",
			nums:     []int{1, 1, 1, 1, 1},
			expected: 4,
		},
		{
			name:     "teleport from start straight to last index",
			nums:     []int{7, 1, 1, 14},
			expected: 1,
		},
		{
			name:     "teleport via prime 3 from index 0 reaches end in one jump",
			nums:     []int{3, 1, 5, 1, 9},
			expected: 1,
		},
		{
			name:     "all composites and ones means no teleport ever fires",
			nums:     []int{1, 4, 6, 8, 9},
			expected: 4,
		},
		{
			name:     "prime appears only once and has no other multiples",
			nums:     []int{1, 7, 1, 1, 1},
			expected: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := minJumps(tt.nums)
			if got != tt.expected {
				t.Errorf("minJumps(%v) = %d, want %d", tt.nums, got, tt.expected)
			}
		})
	}
}
