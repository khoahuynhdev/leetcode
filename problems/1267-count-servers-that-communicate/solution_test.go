package main

import "testing"

func TestCountServers(t *testing.T) {
	tests := []struct {
		name     string
		grid     [][]int
		expected int
	}{
		{"example 1: diagonal servers cannot communicate", [][]int{{1, 0}, {0, 1}}, 0},
		{"example 2: L-shape all communicate", [][]int{{1, 0}, {1, 1}}, 3},
		{"example 3: row pair and column pair plus isolated", [][]int{{1, 1, 0, 0}, {0, 0, 1, 0}, {0, 0, 1, 0}, {0, 0, 0, 1}}, 4},
		{"edge case: single cell with server", [][]int{{1}}, 0},
		{"edge case: single cell without server", [][]int{{0}}, 0},
		{"edge case: no servers", [][]int{{0, 0, 0}, {0, 0, 0}}, 0},
		{"edge case: all servers", [][]int{{1, 1}, {1, 1}}, 4},
		{"edge case: single row with two servers", [][]int{{1, 0, 0, 1, 0}}, 2},
		{"edge case: single column with one server", [][]int{{0}, {1}, {0}}, 0},
		{"edge case: server shares row and column, counted once", [][]int{{1, 1}, {1, 0}}, 3},
		{"edge case: isolated servers in larger grid", [][]int{{1, 0, 0}, {0, 0, 1}, {0, 1, 0}}, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := countServers(tt.grid)
			if result != tt.expected {
				t.Errorf("got %v, want %v", result, tt.expected)
			}
		})
	}
}
