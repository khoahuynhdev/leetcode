package main

import "testing"

func TestSolution(t *testing.T) {
	tests := []struct {
		name      string
		classroom []string
		energy    int
		expected  int
	}{
		{
			name:      "example 1: detour around an obstacle",
			classroom: []string{"S.", "XL"},
			energy:    2,
			expected:  2,
		},
		{
			name:      "example 2: reset area refills energy mid-route",
			classroom: []string{"LS", "RL"},
			energy:    4,
			expected:  3,
		},
		{
			name:      "example 3: no route collects every litter",
			classroom: []string{"L.S", "RXL"},
			energy:    3,
			expected:  -1,
		},
		{
			name:      "edge case: no litter at all needs zero moves",
			classroom: []string{"S..", "..."},
			energy:    1,
			expected:  0,
		},
		{
			name:      "edge case: single cell grid with only the student",
			classroom: []string{"S"},
			energy:    50,
			expected:  0,
		},
		{
			name:      "edge case: litter walled off by an obstacle",
			classroom: []string{"SXL"},
			energy:    50,
			expected:  -1,
		},
		{
			name:      "edge case: energy exactly covers the straight walk",
			classroom: []string{"S.L"},
			energy:    2,
			expected:  2,
		},
		{
			name:      "edge case: one unit short of reaching the litter",
			classroom: []string{"S.L"},
			energy:    1,
			expected:  -1,
		},
		{
			name:      "edge case: reset area makes a long corridor feasible",
			classroom: []string{"S.R.L"},
			energy:    2,
			expected:  4,
		},
		{
			name:      "edge case: backtracking to collect litter on both sides",
			classroom: []string{"LSL"},
			energy:    3,
			expected:  3,
		},
		{
			name:      "edge case: backtracking impossible without enough energy",
			classroom: []string{"LSL"},
			energy:    2,
			expected:  -1,
		},
		{
			name:      "edge case: student starts adjacent to the only litter",
			classroom: []string{"SL"},
			energy:    1,
			expected:  1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := minMoves(tt.classroom, tt.energy)
			if got != tt.expected {
				t.Errorf("minMoves(%v, %d) = %v, want %v", tt.classroom, tt.energy, got, tt.expected)
			}
		})
	}
}
