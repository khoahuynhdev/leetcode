package main

import "testing"

func TestCountIslands(t *testing.T) {
	tests := []struct {
		name     string
		grid     [][]int
		k        int
		expected int
	}{
		{
			"example 1: four islands, two divisible by 5",
			[][]int{{0, 2, 1, 0, 0}, {0, 5, 0, 0, 5}, {0, 0, 1, 0, 0}, {0, 1, 4, 7, 0}, {0, 2, 0, 0, 8}},
			5,
			2,
		},
		{
			"example 2: six diagonal-only single-cell islands",
			[][]int{{3, 0, 3, 0}, {0, 3, 0, 3}, {3, 0, 3, 0}},
			3,
			6,
		},
		{
			"edge case: all water",
			[][]int{{0, 0}, {0, 0}},
			7,
			0,
		},
		{
			"edge case: single cell not divisible",
			[][]int{{4}},
			3,
			0,
		},
		{
			"edge case: k equals 1 counts every island",
			[][]int{{1, 0, 2}, {0, 0, 0}, {3, 0, 4}},
			1,
			4,
		},
		{
			"edge case: whole grid one island",
			[][]int{{1, 2}, {3, 4}},
			10,
			1,
		},
		{
			"edge case: large values sum beyond 32-bit",
			[][]int{{1000000, 1000000, 1000000}, {1000000, 1000000, 1000000}},
			1000000,
			1,
		},
		{
			"edge case: snake-shaped island",
			[][]int{{1, 1, 1}, {0, 0, 1}, {1, 1, 1}},
			7,
			1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := countIslands(tt.grid, tt.k)
			if result != tt.expected {
				t.Errorf("got %v, want %v", result, tt.expected)
			}
		})
	}
}
