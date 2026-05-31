package main

import "testing"

func TestAsteroidsDestroyed(t *testing.T) {
	tests := []struct {
		name      string
		mass      int
		asteroids []int
		expected  bool
	}{
		{
			name:      "example 1: mass 10 vs [3,9,19,5,21] all destroyable",
			mass:      10,
			asteroids: []int{3, 9, 19, 5, 21},
			expected:  true,
		},
		{
			name:      "example 2: mass 5 vs [4,9,23,4] cannot reach 23",
			mass:      5,
			asteroids: []int{4, 9, 23, 4},
			expected:  false,
		},
		{
			name:      "edge case: single asteroid equal to mass (>= rule)",
			mass:      5,
			asteroids: []int{5},
			expected:  true,
		},
		{
			name:      "edge case: single asteroid heavier than mass",
			mass:      1,
			asteroids: []int{2},
			expected:  false,
		},
		{
			name:      "edge case: single asteroid lighter than mass",
			mass:      100,
			asteroids: []int{1},
			expected:  true,
		},
		{
			name:      "edge case: order matters - sorted greedy succeeds",
			mass:      5,
			asteroids: []int{5, 10, 20},
			expected:  true,
		},
		{
			name:      "edge case: all equal to initial mass",
			mass:      3,
			asteroids: []int{3, 3, 3, 3},
			expected:  true,
		},
		{
			name:      "edge case: large cumulative mass (overflow guard)",
			mass:      100000,
			asteroids: makeRepeated(100000, 100000),
			expected:  true,
		},
		{
			name:      "edge case: first asteroid already too heavy",
			mass:      1,
			asteroids: []int{2, 3, 4, 5},
			expected:  false,
		},
		{
			name:      "edge case: minimum mass barely survives chain",
			mass:      1,
			asteroids: []int{1, 2, 4, 8, 16},
			expected:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := append([]int(nil), tt.asteroids...)
			got := asteroidsDestroyed(tt.mass, input)
			if got != tt.expected {
				t.Errorf("asteroidsDestroyed(%d, %v) = %v, want %v",
					tt.mass, tt.asteroids, got, tt.expected)
			}
		})
	}
}

func makeRepeated(value, count int) []int {
	out := make([]int, count)
	for i := range out {
		out[i] = value
	}
	return out
}
