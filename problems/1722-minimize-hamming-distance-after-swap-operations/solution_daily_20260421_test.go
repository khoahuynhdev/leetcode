package main

import "testing"

func TestMinimumHammingDistanceDaily20260421(t *testing.T) {
	tests := []struct {
		name         string
		source       []int
		target       []int
		allowedSwaps [][]int
		expected     int
	}{
		{
			name:         "example 1: two disjoint swap pairs",
			source:       []int{1, 2, 3, 4},
			target:       []int{2, 1, 4, 5},
			allowedSwaps: [][]int{{0, 1}, {2, 3}},
			expected:     1,
		},
		{
			name:         "example 2: no allowed swaps",
			source:       []int{1, 2, 3, 4},
			target:       []int{1, 3, 2, 4},
			allowedSwaps: [][]int{},
			expected:     2,
		},
		{
			name:         "example 3: one big component reaches zero",
			source:       []int{5, 1, 2, 4, 3},
			target:       []int{1, 5, 4, 2, 3},
			allowedSwaps: [][]int{{0, 4}, {4, 2}, {1, 3}, {1, 4}},
			expected:     0,
		},
		{
			name:         "edge case: single element already matching",
			source:       []int{7},
			target:       []int{7},
			allowedSwaps: [][]int{},
			expected:     0,
		},
		{
			name:         "edge case: single element mismatching",
			source:       []int{7},
			target:       []int{8},
			allowedSwaps: [][]int{},
			expected:     1,
		},
		{
			name:         "edge case: already equal with a swap available",
			source:       []int{1, 2, 3},
			target:       []int{1, 2, 3},
			allowedSwaps: [][]int{{0, 1}},
			expected:     0,
		},
		{
			name:         "edge case: duplicate values with multiset intersection",
			source:       []int{1, 2, 1, 2},
			target:       []int{2, 1, 2, 1},
			allowedSwaps: [][]int{{0, 1}, {2, 3}},
			expected:     0,
		},
		{
			name:         "edge case: redundant and reversed edges",
			source:       []int{1, 2, 3, 4},
			target:       []int{4, 3, 2, 1},
			allowedSwaps: [][]int{{0, 1}, {1, 0}, {1, 2}, {2, 3}},
			expected:     0,
		},
		{
			name:         "edge case: component values do not overlap target",
			source:       []int{1, 2},
			target:       []int{3, 4},
			allowedSwaps: [][]int{{0, 1}},
			expected:     2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := minimumHammingDistanceDaily20260421(tt.source, tt.target, tt.allowedSwaps)
			if got != tt.expected {
				t.Errorf("minimumHammingDistanceDaily20260421(%v, %v, %v) = %d, want %d",
					tt.source, tt.target, tt.allowedSwaps, got, tt.expected)
			}
		})
	}
}
