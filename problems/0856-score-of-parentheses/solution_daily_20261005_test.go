package main

import "testing"

func TestSolution(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		expected int
	}{
		{"example 1: single pair", "()", 1},
		{"example 2: nested pair", "(())", 2},
		{"example 3: concatenated pairs", "()()", 2},
		{"edge case: deep nesting", "((((()))))", 16},
		{"edge case: many concatenated pairs", "()()()()", 4},
		{"edge case: mixed nesting and concatenation", "(()(()))", 6},
		{"edge case: nested group followed by pair", "(()())()", 5},
		{"edge case: maximum depth for length 50", "(((((((((((((((((((((((((" + "))))))))))))))))))))))))))", 1 << 24},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := scoreOfParentheses(tt.s)
			if result != tt.expected {
				t.Errorf("got %v, want %v", result, tt.expected)
			}
		})
	}
}
