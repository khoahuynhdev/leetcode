package main

import "testing"

func TestReverseParentheses(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		expected string
	}{
		{"example 1: single pair", "(abcd)", "dcba"},
		{"example 2: two nested levels", "(u(love)i)", "iloveu"},
		{"example 3: three nested levels", "(ed(et(oc))el)", "leetcode"},
		{"edge case: no parentheses", "abc", "abc"},
		{"edge case: empty parentheses", "a()b", "ab"},
		{"edge case: only empty parentheses", "()", ""},
		{"edge case: double nesting cancels", "((ab))", "ab"},
		{"edge case: adjacent groups", "(ab)(cd)", "badc"},
		{"edge case: letters outside brackets", "a(bc)d", "acbd"},
		{"edge case: single character", "a", "a"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := reverseParentheses(tt.s)
			if result != tt.expected {
				t.Errorf("got %q, want %q", result, tt.expected)
			}
		})
	}
}
