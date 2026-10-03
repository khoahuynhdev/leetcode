package main

import "testing"

func TestLongestValidParentheses(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		expected int
	}{
		{"example 1: unmatched leading open", "(()", 2},
		{"example 2: unmatched on both ends", ")()())", 4},
		{"example 3: empty string", "", 0},
		{"edge case: single open", "(", 0},
		{"edge case: single close", ")", 0},
		{"edge case: all opens", "((((", 0},
		{"edge case: all closes", "))))", 0},
		{"edge case: wrong order", ")(", 0},
		{"edge case: entire string valid nested and adjacent", "()(())", 6},
		{"edge case: deeply nested", "((()))", 6},
		{"edge case: break splits runs", "()(()", 2},
		{"edge case: longer run after break", "())(())()", 6},
		{"edge case: unmatched open in middle", "(()(((()", 2},
		{"edge case: valid run at end", "))((()))", 6},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := longestValidParentheses(tt.s)
			if result != tt.expected {
				t.Errorf("longestValidParentheses(%q) = %d, want %d", tt.s, result, tt.expected)
			}
		})
	}
}
