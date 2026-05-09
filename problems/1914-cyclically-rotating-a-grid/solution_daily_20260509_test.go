package main

import (
	"reflect"
	"testing"
)

func TestRotateGridDaily20260509(t *testing.T) {
	tests := []struct {
		name     string
		grid     [][]int
		k        int
		expected [][]int
	}{
		{
			name:     "example 1: 2x2 grid with k=1",
			grid:     [][]int{{40, 10}, {30, 20}},
			k:        1,
			expected: [][]int{{10, 20}, {40, 30}},
		},
		{
			name:     "example 2: 4x4 grid with k=2",
			grid:     [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}},
			k:        2,
			expected: [][]int{{3, 4, 8, 12}, {2, 11, 10, 16}, {1, 7, 6, 15}, {5, 9, 13, 14}},
		},
		{
			name:     "edge case: k equals perimeter so 2x2 is unchanged",
			grid:     [][]int{{1, 2}, {3, 4}},
			k:        4,
			expected: [][]int{{1, 2}, {3, 4}},
		},
		{
			name:     "edge case: k larger than perimeter wraps via modulo (2x2, k=5 acts like k=1)",
			grid:     [][]int{{40, 10}, {30, 20}},
			k:        5,
			expected: [][]int{{10, 20}, {40, 30}},
		},
		{
			name:     "edge case: thin 2x4 grid (single layer, rectangular ring)",
			grid:     [][]int{{1, 2, 3, 4}, {8, 7, 6, 5}},
			k:        1,
			expected: [][]int{{2, 3, 4, 5}, {1, 8, 7, 6}},
		},
		{
			name:     "edge case: thin 4x2 grid (single layer, tall ring)",
			grid:     [][]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}},
			k:        1,
			expected: [][]int{{2, 4}, {1, 6}, {3, 8}, {5, 7}},
		},
		{
			name:     "edge case: huge k=1000000000 on 4x4 reduces per-layer (outer perim 12 → 1000000000%12=4, inner perim 4 → 0)",
			grid:     [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}},
			k:        1000000000,
			expected: [][]int{{8, 12, 16, 15}, {4, 6, 7, 14}, {3, 10, 11, 13}, {2, 1, 5, 9}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// rotateGrid mutates the grid in place; clone the input so each
			// table entry is independent and the test name reflects the
			// original input.
			input := make([][]int, len(tt.grid))
			for i := range tt.grid {
				input[i] = append([]int(nil), tt.grid[i]...)
			}

			got := rotateGrid(input, tt.k)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("rotateGrid(%v, %d)\n got = %v\nwant = %v", tt.grid, tt.k, got, tt.expected)
			}
		})
	}
}
