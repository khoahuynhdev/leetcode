package main

import "testing"

func TestIsRectangleOverlap(t *testing.T) {
	tests := []struct {
		name     string
		rec1     []int
		rec2     []int
		expected bool
	}{
		{
			name:     "example 1: partially overlapping corner to corner",
			rec1:     []int{0, 0, 2, 2},
			rec2:     []int{1, 1, 3, 3},
			expected: true,
		},
		{
			name:     "example 2: sharing a vertical edge only",
			rec1:     []int{0, 0, 1, 1},
			rec2:     []int{1, 0, 2, 1},
			expected: false,
		},
		{
			name:     "example 3: diagonally separated, no contact",
			rec1:     []int{0, 0, 1, 1},
			rec2:     []int{2, 2, 3, 3},
			expected: false,
		},
		{
			name:     "edge case: touching at exactly one corner",
			rec1:     []int{0, 0, 1, 1},
			rec2:     []int{1, 1, 2, 2},
			expected: false,
		},
		{
			name:     "edge case: rec2 fully nested inside rec1",
			rec1:     []int{0, 0, 10, 10},
			rec2:     []int{2, 2, 3, 3},
			expected: true,
		},
		{
			name:     "edge case: rec1 fully nested inside rec2",
			rec1:     []int{4, 4, 5, 5},
			rec2:     []int{-1, -1, 9, 9},
			expected: true,
		},
		{
			name:     "edge case: identical rectangles overlap completely",
			rec1:     []int{-3, -3, 3, 3},
			rec2:     []int{-3, -3, 3, 3},
			expected: true,
		},
		{
			name:     "edge case: sharing a horizontal edge only",
			rec1:     []int{0, 0, 2, 1},
			rec2:     []int{0, 1, 2, 2},
			expected: false,
		},
		{
			name:     "edge case: x overlaps but y is disjoint",
			rec1:     []int{0, 0, 5, 1},
			rec2:     []int{1, 4, 4, 6},
			expected: false,
		},
		{
			name:     "edge case: y overlaps but x is disjoint",
			rec1:     []int{0, 0, 1, 5},
			rec2:     []int{3, 1, 6, 4},
			expected: false,
		},
		{
			name:     "edge case: crossing plus shape overlaps",
			rec1:     []int{-5, -1, 5, 1},
			rec2:     []int{-1, -5, 1, 5},
			expected: true,
		},
		{
			name:     "edge case: entirely negative coordinates overlapping",
			rec1:     []int{-9, -9, -5, -5},
			rec2:     []int{-7, -7, -1, -1},
			expected: true,
		},
		{
			name:     "edge case: extreme coordinate bounds overlapping",
			rec1:     []int{-1000000000, -1000000000, 0, 0},
			rec2:     []int{-1, -1, 1000000000, 1000000000},
			expected: true,
		},
		{
			name:     "edge case: extreme coordinate bounds touching at origin",
			rec1:     []int{-1000000000, -1000000000, 0, 0},
			rec2:     []int{0, 0, 1000000000, 1000000000},
			expected: false,
		},
		{
			name:     "edge case: overlap is a thin sliver of positive area",
			rec1:     []int{0, 0, 10, 10},
			rec2:     []int{9, -5, 20, 15},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isRectangleOverlap(tt.rec1, tt.rec2); got != tt.expected {
				t.Errorf("isRectangleOverlap(%v, %v) = %v, want %v", tt.rec1, tt.rec2, got, tt.expected)
			}

			// Overlap is symmetric: swapping the arguments must not change the result.
			if got := isRectangleOverlap(tt.rec2, tt.rec1); got != tt.expected {
				t.Errorf("isRectangleOverlap(%v, %v) = %v, want %v (symmetry)", tt.rec2, tt.rec1, got, tt.expected)
			}
		})
	}
}
