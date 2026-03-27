package main

import "testing"

func TestAreSimilar(t *testing.T) {
	tests := []struct {
		name     string
		mat      [][]int
		k        int
		expected bool
	}{
		{
			name:     "example 1: 3x3 matrix with k=4",
			mat:      [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			k:        4,
			expected: false,
		},
		{
			name:     "example 2: repeating pattern with k=2",
			mat:      [][]int{{1, 2, 1, 2}, {5, 5, 5, 5}, {6, 3, 6, 3}},
			k:        2,
			expected: true,
		},
		{
			name:     "example 3: all same elements",
			mat:      [][]int{{2, 2}, {2, 2}},
			k:        3,
			expected: true,
		},
		{
			name:     "edge case: k is multiple of n",
			mat:      [][]int{{1, 2, 3}, {4, 5, 6}},
			k:        6,
			expected: true,
		},
		{
			name:     "edge case: single column matrix",
			mat:      [][]int{{1}, {2}, {3}},
			k:        50,
			expected: true,
		},
		{
			name:     "edge case: single row single column",
			mat:      [][]int{{7}},
			k:        1,
			expected: true,
		},
		{
			name:     "edge case: k=1 with non-uniform row",
			mat:      [][]int{{1, 2}},
			k:        1,
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := areSimilar(tt.mat, tt.k)
			if result != tt.expected {
				t.Errorf("areSimilar(%v, %d) = %v, want %v", tt.mat, tt.k, result, tt.expected)
			}
		})
	}
}
