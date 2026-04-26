package main

import "testing"

func TestContainsCycleDaily20260426(t *testing.T) {
	tests := []struct {
		name     string
		grid     [][]byte
		expected bool
	}{
		{
			name: "example 1: 4x4 outer ring of 'a' forms a cycle",
			grid: [][]byte{
				{'a', 'a', 'a', 'a'},
				{'a', 'b', 'b', 'a'},
				{'a', 'b', 'b', 'a'},
				{'a', 'a', 'a', 'a'},
			},
			expected: true,
		},
		{
			name: "example 2: 4x4 with embedded c-shaped cycle",
			grid: [][]byte{
				{'c', 'c', 'c', 'a'},
				{'c', 'd', 'c', 'c'},
				{'c', 'c', 'e', 'c'},
				{'f', 'c', 'c', 'c'},
			},
			expected: true,
		},
		{
			name: "example 3: 3x3 no cycle",
			grid: [][]byte{
				{'a', 'b', 'b'},
				{'b', 'z', 'b'},
				{'b', 'b', 'a'},
			},
			expected: false,
		},
		{
			name:     "edge case: single cell cannot form a cycle",
			grid:     [][]byte{{'a'}},
			expected: false,
		},
		{
			name:     "edge case: single row of same char (no cycle in 1xN)",
			grid:     [][]byte{{'a', 'a', 'a', 'a'}},
			expected: false,
		},
		{
			name: "edge case: single column of same char (no cycle in Mx1)",
			grid: [][]byte{
				{'a'},
				{'a'},
				{'a'},
				{'a'},
			},
			expected: false,
		},
		{
			name: "edge case: smallest possible cycle (2x2 of same char)",
			grid: [][]byte{
				{'a', 'a'},
				{'a', 'a'},
			},
			expected: true,
		},
		{
			name: "edge case: 2x2 with three matching cells (no cycle)",
			grid: [][]byte{
				{'a', 'a'},
				{'b', 'a'},
			},
			expected: false,
		},
		{
			name: "edge case: all distinct characters (no same-value edges)",
			grid: [][]byte{
				{'a', 'b', 'c'},
				{'d', 'e', 'f'},
				{'g', 'h', 'i'},
			},
			expected: false,
		},
		{
			name: "edge case: U-shaped path of same char has no cycle",
			grid: [][]byte{
				{'a', 'a', 'a'},
				{'b', 'b', 'a'},
				{'a', 'a', 'a'},
			},
			expected: false,
		},
		{
			name: "edge case: cycle exists only in second component",
			grid: [][]byte{
				{'a', 'b', 'b'},
				{'c', 'b', 'b'},
				{'d', 'e', 'f'},
			},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := containsCycleDaily20260426(tt.grid)
			if got != tt.expected {
				t.Errorf("containsCycleDaily20260426(...) = %v, want %v", got, tt.expected)
			}
		})
	}
}
