package main

import "testing"

func toGrid(rows []string) [][]byte {
	grid := make([][]byte, len(rows))
	for i, r := range rows {
		grid[i] = []byte(r)
	}
	return grid
}

func TestHasValidPath(t *testing.T) {
	tests := []struct {
		name     string
		grid     []string
		expected bool
	}{
		{"example 1: multiple valid paths exist", []string{"(((", ")()", "(()", "(()"}, true},
		{"example 2: both paths start with ')'", []string{"))", "(("}, false},
		{"edge case: single cell has odd length", []string{"("}, false},
		{"edge case: 1x2 simple pair", []string{"()"}, true},
		{"edge case: 1x2 reversed pair", []string{")("}, false},
		{"edge case: odd path length 2x2", []string{"()", ")("}, false},
		{"edge case: single column valid", []string{"(", "(", ")", ")"}, true},
		{"edge case: last cell is '('", []string{"((", ")("}, false},
		{"edge case: 3x3 grid has odd path length", []string{"(()", ")()", "())"}, false},
		{"edge case: 2x3 valid path", []string{"(()", "())"}, true},
		{"edge case: 2x3 even length and good endpoints but no valid path", []string{"())", ")))"}, false},
		{"edge case: 3x4 valid path requires correct turns", []string{"(()(", ")(()", "()))"}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := hasValidPath(toGrid(tt.grid))
			if result != tt.expected {
				t.Errorf("got %v, want %v", result, tt.expected)
			}
		})
	}
}
