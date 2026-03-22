package main

import "testing"

func TestFindRotation(t *testing.T) {
	tests := []struct {
		name     string
		mat      [][]int
		target   [][]int
		expected bool
	}{
		{
			name:     "example 1: 90 degree rotation",
			mat:      [][]int{{0, 1}, {1, 0}},
			target:   [][]int{{1, 0}, {0, 1}},
			expected: true,
		},
		{
			name:     "example 2: no rotation matches",
			mat:      [][]int{{0, 1}, {1, 1}},
			target:   [][]int{{1, 0}, {0, 1}},
			expected: false,
		},
		{
			name:     "example 3: 180 degree rotation",
			mat:      [][]int{{0, 0, 0}, {0, 1, 0}, {1, 1, 1}},
			target:   [][]int{{1, 1, 1}, {0, 1, 0}, {0, 0, 0}},
			expected: true,
		},
		{
			name:     "edge case: 1x1 same value",
			mat:      [][]int{{1}},
			target:   [][]int{{1}},
			expected: true,
		},
		{
			name:     "edge case: 1x1 different value",
			mat:      [][]int{{0}},
			target:   [][]int{{1}},
			expected: false,
		},
		{
			name:     "edge case: already equal no rotation needed",
			mat:      [][]int{{1, 0}, {0, 1}},
			target:   [][]int{{1, 0}, {0, 1}},
			expected: true,
		},
		{
			name:     "edge case: 270 degree rotation",
			mat:      [][]int{{0, 1}, {1, 0}},
			target:   [][]int{{0, 1}, {1, 0}},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := findRotation(tt.mat, tt.target)
			if result != tt.expected {
				t.Errorf("got %v, want %v", result, tt.expected)
			}
		})
	}
}
