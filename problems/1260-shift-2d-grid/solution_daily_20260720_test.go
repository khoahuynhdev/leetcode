package main

import (
	"reflect"
	"testing"
)

func TestShiftGrid(t *testing.T) {
	tests := []struct {
		name     string
		grid     [][]int
		k        int
		expected [][]int
	}{
		{
			name:     "example 1: shift 3x3 by 1",
			grid:     [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			k:        1,
			expected: [][]int{{9, 1, 2}, {3, 4, 5}, {6, 7, 8}},
		},
		{
			name:     "example 2: shift 4x4 by 4",
			grid:     [][]int{{3, 8, 1, 9}, {19, 7, 2, 5}, {4, 6, 11, 10}, {12, 0, 21, 13}},
			k:        4,
			expected: [][]int{{12, 0, 21, 13}, {3, 8, 1, 9}, {19, 7, 2, 5}, {4, 6, 11, 10}},
		},
		{
			name:     "example 3: full rotation is a no-op (k equals m*n)",
			grid:     [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			k:        9,
			expected: [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
		},
		{
			name:     "edge case: k is 0 leaves grid unchanged",
			grid:     [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			k:        0,
			expected: [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
		},
		{
			name:     "edge case: k larger than m*n reduces via modulo",
			grid:     [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			k:        10,
			expected: [][]int{{9, 1, 2}, {3, 4, 5}, {6, 7, 8}},
		},
		{
			name:     "edge case: single row behaves like 1D cyclic shift",
			grid:     [][]int{{1, 2, 3, 4}},
			k:        2,
			expected: [][]int{{3, 4, 1, 2}},
		},
		{
			name:     "edge case: single column wraps bottom to top",
			grid:     [][]int{{1}, {2}, {3}},
			k:        1,
			expected: [][]int{{3}, {1}, {2}},
		},
		{
			name:     "edge case: 1x1 grid stays put for any k",
			grid:     [][]int{{7}},
			k:        100,
			expected: [][]int{{7}},
		},
		{
			name:     "edge case: negative values are preserved",
			grid:     [][]int{{-1, -2}, {-3, -4}},
			k:        1,
			expected: [][]int{{-4, -1}, {-2, -3}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := shiftGrid(tt.grid, tt.k)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("shiftGrid(%v, %d) = %v, want %v", tt.grid, tt.k, result, tt.expected)
			}
		})
	}
}
