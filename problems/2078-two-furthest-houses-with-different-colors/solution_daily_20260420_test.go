package main

import "testing"

func TestMaxDistance(t *testing.T) {
	tests := []struct {
		name     string
		colors   []int
		expected int
	}{
		{
			name:     "example 1: [1,1,1,6,1,1,1] -> 3",
			colors:   []int{1, 1, 1, 6, 1, 1, 1},
			expected: 3,
		},
		{
			name:     "example 2: [1,8,3,8,3] -> 4",
			colors:   []int{1, 8, 3, 8, 3},
			expected: 4,
		},
		{
			name:     "example 3: [0,1] -> 1",
			colors:   []int{0, 1},
			expected: 1,
		},
		{
			name:     "edge case: minimum length same-prefix [5,5,5,7] -> 3",
			colors:   []int{5, 5, 5, 7},
			expected: 3,
		},
		{
			name:     "edge case: odd color at left endpoint [6,1,1,1,1] -> 4",
			colors:   []int{6, 1, 1, 1, 1},
			expected: 4,
		},
		{
			name:     "edge case: odd color stuck in the middle [1,1,6,1,1] -> 2",
			colors:   []int{1, 1, 6, 1, 1},
			expected: 2,
		},
		{
			name:     "edge case: first and last already differ [0,0,0,0,100] -> 4",
			colors:   []int{0, 0, 0, 0, 100},
			expected: 4,
		},
		{
			name:     "edge case: two houses with different colors [2,7] -> 1",
			colors:   []int{2, 7},
			expected: 1,
		},
		{
			name:     "edge case: alternating colors [1,2,1,2,1,2] -> 5",
			colors:   []int{1, 2, 1, 2, 1, 2},
			expected: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxDistance(tt.colors)
			if got != tt.expected {
				t.Errorf("maxDistance(%v) = %d, want %d", tt.colors, got, tt.expected)
			}
		})
	}
}
