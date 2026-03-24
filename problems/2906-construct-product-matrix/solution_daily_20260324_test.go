package main

import (
	"reflect"
	"testing"
)

func TestConstructProductMatrix(t *testing.T) {
	tests := []struct {
		name     string
		grid     [][]int
		expected [][]int
	}{
		{
			name:     "example 1: 2x2 grid",
			grid:     [][]int{{1, 2}, {3, 4}},
			expected: [][]int{{24, 12}, {8, 6}},
		},
		{
			name:     "example 2: single column with 12345",
			grid:     [][]int{{12345}, {2}, {1}},
			expected: [][]int{{2}, {0}, {0}},
		},
		{
			name:     "edge case: minimum size grid (2 elements)",
			grid:     [][]int{{5, 7}},
			expected: [][]int{{7, 5}},
		},
		{
			name:     "edge case: all elements are 1",
			grid:     [][]int{{1, 1, 1}, {1, 1, 1}},
			expected: [][]int{{1, 1, 1}, {1, 1, 1}},
		},
		{
			name:     "edge case: multiple multiples of 12345",
			grid:     [][]int{{12345, 12345}, {2, 3}},
			expected: [][]int{{0, 0}, {0, 0}},
		},
		{
			name:     "edge case: large values",
			grid:     [][]int{{1000000000, 1000000000}},
			expected: [][]int{{5620, 5620}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := constructProductMatrix(tt.grid)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("got %v, want %v", result, tt.expected)
			}
		})
	}
}
