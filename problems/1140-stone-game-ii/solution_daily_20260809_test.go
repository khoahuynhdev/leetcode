package main

import "testing"

func TestSolution(t *testing.T) {
	tests := []struct {
		name     string
		piles    []int
		expected int
	}{
		{
			name:     "example 1: alice takes one pile first to leave a bad split for bob",
			piles:    []int{2, 7, 9, 4, 4},
			expected: 10,
		},
		{
			name:     "example 2: alice waits for the huge last pile",
			piles:    []int{1, 2, 3, 4, 5, 100},
			expected: 104,
		},
		{
			name:     "edge case: single pile is taken entirely by alice",
			piles:    []int{7},
			expected: 7,
		},
		{
			name:     "edge case: two piles fit in alice's first move (X <= 2M = 2)",
			piles:    []int{5, 3},
			expected: 8,
		},
		{
			name:     "edge case: empty input",
			piles:    []int{},
			expected: 0,
		},
		{
			name:     "edge case: all piles equal, even count splits evenly",
			piles:    []int{1, 1, 1, 1},
			expected: 2,
		},
		{
			name:     "edge case: three piles, alice must take two or bob sweeps the tail",
			piles:    []int{3, 1, 2},
			expected: 4,
		},
		{
			name:     "edge case: strictly increasing piles reward waiting",
			piles:    []int{1, 2, 3, 4, 5, 6, 7},
			expected: 15,
		},
		{
			name:     "edge case: max-value piles at the constraint bound",
			piles:    []int{10000, 10000, 10000, 10000, 10000},
			expected: 30000,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := stoneGameII(tt.piles); got != tt.expected {
				t.Errorf("stoneGameII(%v) = %v, want %v", tt.piles, got, tt.expected)
			}
		})
	}
}
