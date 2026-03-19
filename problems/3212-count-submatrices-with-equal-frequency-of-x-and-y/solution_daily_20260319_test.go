package main

import "testing"

func TestCountSubmatrices(t *testing.T) {
	tests := []struct {
		name     string
		grid     [][]byte
		expected int
	}{
		{
			name: "example 1: X Y dot grid",
			grid: [][]byte{
				{'X', 'Y', '.'},
				{'Y', '.', '.'},
			},
			expected: 3,
		},
		{
			name: "example 2: more X than Y",
			grid: [][]byte{
				{'X', 'X'},
				{'X', 'Y'},
			},
			expected: 0,
		},
		{
			name: "example 3: all dots",
			grid: [][]byte{
				{'.', '.'},
				{'.', '.'},
			},
			expected: 0,
		},
		{
			name:     "edge case: single X",
			grid:     [][]byte{{'X'}},
			expected: 0,
		},
		{
			name:     "edge case: single dot",
			grid:     [][]byte{{'.'}},
			expected: 0,
		},
		{
			name: "edge case: 1x2 with X then Y",
			grid: [][]byte{
				{'X', 'Y'},
			},
			expected: 1,
		},
		{
			name: "edge case: no X in grid",
			grid: [][]byte{
				{'Y', '.'},
				{'.', 'Y'},
			},
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := countSubmatrices(tt.grid)
			if result != tt.expected {
				t.Errorf("got %d, want %d", result, tt.expected)
			}
		})
	}
}
