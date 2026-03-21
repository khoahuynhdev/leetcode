package main

import (
	"reflect"
	"testing"
)

func TestFlipSubmatrix(t *testing.T) {
	tests := []struct {
		name     string
		grid     [][]int
		x        int
		y        int
		k        int
		expected [][]int
	}{
		{
			name:     "example 1: 3x3 submatrix in 4x4 grid",
			grid:     [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}},
			x:        1,
			y:        0,
			k:        3,
			expected: [][]int{{1, 2, 3, 4}, {13, 14, 15, 8}, {9, 10, 11, 12}, {5, 6, 7, 16}},
		},
		{
			name:     "example 2: 2x2 submatrix at top-right of 2x4 grid",
			grid:     [][]int{{3, 4, 2, 3}, {2, 3, 4, 2}},
			x:        0,
			y:        2,
			k:        2,
			expected: [][]int{{3, 4, 4, 2}, {2, 3, 2, 3}},
		},
		{
			name:     "edge case: k=1 single cell no change",
			grid:     [][]int{{5, 6}, {7, 8}},
			x:        0,
			y:        1,
			k:        1,
			expected: [][]int{{5, 6}, {7, 8}},
		},
		{
			name:     "edge case: entire 3x3 grid flipped",
			grid:     [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			x:        0,
			y:        0,
			k:        3,
			expected: [][]int{{7, 8, 9}, {4, 5, 6}, {1, 2, 3}},
		},
		{
			name:     "edge case: 1x1 grid",
			grid:     [][]int{{42}},
			x:        0,
			y:        0,
			k:        1,
			expected: [][]int{{42}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := flipSubmatrix(tt.grid, tt.x, tt.y, tt.k)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("got %v, want %v", result, tt.expected)
			}
		})
	}
}
