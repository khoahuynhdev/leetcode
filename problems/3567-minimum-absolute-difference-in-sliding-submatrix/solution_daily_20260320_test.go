package main

import (
	"reflect"
	"testing"
)

func TestMinimumAbsoluteDifference(t *testing.T) {
	tests := []struct {
		name     string
		grid     [][]int
		k        int
		expected [][]int
	}{
		{
			name:     "example 1: 2x2 grid with k=2",
			grid:     [][]int{{1, 8}, {3, -2}},
			k:        2,
			expected: [][]int{{2}},
		},
		{
			name:     "example 2: 1x2 grid with k=1",
			grid:     [][]int{{3, -1}},
			k:        1,
			expected: [][]int{{0, 0}},
		},
		{
			name:     "example 3: 2x3 grid with k=2",
			grid:     [][]int{{1, -2, 3}, {2, 3, 5}},
			k:        2,
			expected: [][]int{{1, 2}},
		},
		{
			name:     "edge case: 1x1 grid with k=1",
			grid:     [][]int{{42}},
			k:        1,
			expected: [][]int{{0}},
		},
		{
			name:     "edge case: all elements identical",
			grid:     [][]int{{5, 5}, {5, 5}},
			k:        2,
			expected: [][]int{{0}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := minimumAbsoluteDifference(tt.grid, tt.k)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("got %v, want %v", result, tt.expected)
			}
		})
	}
}
