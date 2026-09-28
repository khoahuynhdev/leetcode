package main

import "testing"

func TestMaxDepth(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		expected int
	}{
		{"example 1: mixed expression with depth 3", "(1+(2*3)+((8)/4))+1", 3},
		{"example 2: increasing nested groups", "(1)+((2))+(((3)))", 3},
		{"example 3: parentheses only", "()(())((()()))", 3},
		{"edge case: single digit, no parentheses", "8", 0},
		{"edge case: expression without parentheses", "1+2*3", 0},
		{"edge case: siblings stay at depth 1", "()()()", 1},
		{"edge case: fully nested", "((((()))))", 5},
		{"edge case: deepest nesting at the end", "1+(2+(3-(4/5)))", 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := maxDepth(tt.s)
			if result != tt.expected {
				t.Errorf("maxDepth(%q) = %d, want %d", tt.s, result, tt.expected)
			}
		})
	}
}
