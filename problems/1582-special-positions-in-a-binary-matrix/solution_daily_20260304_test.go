package main

import "testing"

func TestNumSpecial(t *testing.T) {
	tests := []struct {
		name     string
		mat      [][]int
		expected int
	}{
		{
			name:     "example 1: one special position",
			mat:      [][]int{{1, 0, 0}, {0, 0, 1}, {1, 0, 0}},
			expected: 1,
		},
		{
			name:     "example 2: diagonal identity matrix",
			mat:      [][]int{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}},
			expected: 3,
		},
		{
			name:     "edge case: single cell with 1",
			mat:      [][]int{{1}},
			expected: 1,
		},
		{
			name:     "edge case: single cell with 0",
			mat:      [][]int{{0}},
			expected: 0,
		},
		{
			name:     "edge case: all zeros",
			mat:      [][]int{{0, 0, 0}, {0, 0, 0}},
			expected: 0,
		},
		{
			name:     "edge case: all ones",
			mat:      [][]int{{1, 1}, {1, 1}},
			expected: 0,
		},
		{
			name:     "single row with one 1",
			mat:      [][]int{{0, 0, 1, 0}},
			expected: 1,
		},
		{
			name:     "single column with one 1",
			mat:      [][]int{{0}, {1}, {0}},
			expected: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := numSpecial(tt.mat)
			if result != tt.expected {
				t.Errorf("got %d, want %d", result, tt.expected)
			}
		})
	}
}
