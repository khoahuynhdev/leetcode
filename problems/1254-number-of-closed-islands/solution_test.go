package main

import "testing"

func TestClosedIsland(t *testing.T) {
	tests := []struct {
		name     string
		grid     [][]int
		expected int
	}{
		{
			"example 1: two closed islands, one open island on right edge",
			[][]int{
				{1, 1, 1, 1, 1, 1, 1, 0},
				{1, 0, 0, 0, 0, 1, 1, 0},
				{1, 0, 1, 0, 1, 1, 1, 0},
				{1, 0, 0, 0, 0, 1, 0, 1},
				{1, 1, 1, 1, 1, 1, 1, 0},
			},
			2,
		},
		{
			"example 2: single closed cell between border islands",
			[][]int{
				{0, 0, 1, 0, 0},
				{0, 1, 0, 1, 0},
				{0, 1, 1, 1, 0},
			},
			1,
		},
		{
			"example 3: nested closed islands",
			[][]int{
				{1, 1, 1, 1, 1, 1, 1},
				{1, 0, 0, 0, 0, 0, 1},
				{1, 0, 1, 1, 1, 0, 1},
				{1, 0, 1, 0, 1, 0, 1},
				{1, 0, 1, 1, 1, 0, 1},
				{1, 0, 0, 0, 0, 0, 1},
				{1, 1, 1, 1, 1, 1, 1},
			},
			2,
		},
		{
			"edge case: single cell land",
			[][]int{{0}},
			0,
		},
		{
			"edge case: single row",
			[][]int{{1, 0, 1, 0, 1}},
			0,
		},
		{
			"edge case: all water",
			[][]int{
				{1, 1, 1},
				{1, 1, 1},
				{1, 1, 1},
			},
			0,
		},
		{
			"edge case: all land",
			[][]int{
				{0, 0, 0},
				{0, 0, 0},
				{0, 0, 0},
			},
			0,
		},
		{
			"edge case: island leaks to border through one cell",
			[][]int{
				{1, 1, 1, 1, 1},
				{1, 0, 0, 0, 1},
				{1, 0, 1, 0, 0},
				{1, 1, 1, 1, 1},
			},
			0,
		},
		{
			"edge case: diagonal land cells are separate islands",
			[][]int{
				{1, 1, 1, 1},
				{1, 0, 1, 1},
				{1, 1, 0, 1},
				{1, 1, 1, 1},
			},
			2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			grid := make([][]int, len(tt.grid))
			for i, row := range tt.grid {
				grid[i] = append([]int(nil), row...)
			}
			result := closedIsland(grid)
			if result != tt.expected {
				t.Errorf("got %v, want %v", result, tt.expected)
			}
		})
	}
}
