package main

import "testing"

func TestSolution(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{"example 1: one extra closing", "())", 1},
		{"example 2: all opening", "(((", 3},
		{"edge case: already valid", "()", 0},
		{"edge case: nested valid", "(())", 0},
		{"edge case: all closing", ")))", 3},
		{"edge case: reversed pair", ")(", 2},
		{"edge case: single open", "(", 1},
		{"edge case: single close", ")", 1},
		{"edge case: mixed unmatched both sides", "()))((", 4},
		{"edge case: unmatched in middle", "())(", 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := minAddToMakeValid(tt.input)
			if result != tt.expected {
				t.Errorf("minAddToMakeValid(%q) = %d, want %d", tt.input, result, tt.expected)
			}
		})
	}
}
