package main

import "testing"

func TestSolution(t *testing.T) {
	tests := []struct {
		name         string
		n            int
		restrictions [][]int
		expected     int
	}{
		{
			name:         "example 1: caps at buildings 2 and 4",
			n:            5,
			restrictions: [][]int{{2, 1}, {4, 1}},
			expected:     2,
		},
		{
			name:         "example 2: no restrictions, climb freely",
			n:            6,
			restrictions: [][]int{},
			expected:     5,
		},
		{
			name:         "example 3: unsorted restrictions with cap at n",
			n:            10,
			restrictions: [][]int{{5, 3}, {2, 5}, {7, 4}, {10, 3}},
			expected:     5,
		},
		{
			name:         "edge case: smallest n with no restrictions",
			n:            2,
			restrictions: [][]int{},
			expected:     1,
		},
		{
			name:         "edge case: building n hard-capped at zero",
			n:            5,
			restrictions: [][]int{{5, 0}},
			expected:     2,
		},
		{
			name:         "edge case: generous cap tightened by propagation",
			n:            5,
			restrictions: [][]int{{3, 0}},
			expected:     2,
		},
		{
			name:         "edge case: large n unrestricted",
			n:            1000000000,
			restrictions: [][]int{},
			expected:     999999999,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := maxBuilding(tt.n, tt.restrictions)
			if result != tt.expected {
				t.Errorf("maxBuilding(%d, %v) = %d, want %d", tt.n, tt.restrictions, result, tt.expected)
			}
		})
	}
}
