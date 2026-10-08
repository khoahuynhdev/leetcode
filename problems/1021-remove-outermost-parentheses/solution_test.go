package main

import "testing"

func TestSolution(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"example 1: two primitives", "(()())(())", "()()()"},
		{"example 2: three primitives", "(()())(())(()(()))", "()()()()(())"},
		{"example 3: flat primitives become empty", "()()", ""},
		{"edge case: single pair", "()", ""},
		{"edge case: deeply nested single primitive", "((()))", "(())"},
		{"edge case: mixed nested and flat primitives", "()(())()((()))", "()(())"},
		{"edge case: siblings inside one primitive", "(()()())", "()()()"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := removeOuterParentheses(tt.input)
			if result != tt.expected {
				t.Errorf("got %q, want %q", result, tt.expected)
			}
		})
	}
}
