package main

import "testing"

func TestMinCost(t *testing.T) {
	tests := []struct {
		name     string
		grid     [][]int
		expected int
	}{
		{"example 1: 2x2 grid, down-first path wins", [][]int{{1, 2}, {3, 4}}, 6},
		{"example 2: 2x2 grid, right-first path wins", [][]int{{6, 7}, {5, 8}}, 9},
		{"example 3: single row cancels to zero", [][]int{{2, 7, 5}}, 0},
		{"edge case: single cell", [][]int{{5}}, 5},
		{"edge case: single cell zero", [][]int{{0}}, 0},
		{"edge case: single column", [][]int{{1}, {2}, {4}}, 7},
		{"edge case: greedy min prefix fails", [][]int{{0, 1}, {2, 3}}, 1},
		{"edge case: max values", [][]int{{1023, 1023}, {1023, 1023}}, 1023},
		{"edge case: 3x3 all zeros", [][]int{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}}, 0},
		{"edge case: 3x3 mixed", [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := minCost(tt.grid)
			if result != tt.expected {
				t.Errorf("got %v, want %v", result, tt.expected)
			}
		})
	}
}
