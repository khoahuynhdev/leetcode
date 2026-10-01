package main

import "testing"

func TestSolution(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{"example 1: single pair", "()", true},
		{"example 2: consecutive pairs of each type", "()[]{}", true},
		{"example 3: mismatched types", "(]", false},
		{"example 4: nested pairs", "([])", true},
		{"example 5: interleaved pairs", "([)]", false},
		{"edge case: single opening bracket", "(", false},
		{"edge case: single closing bracket", "]", false},
		{"edge case: only closing brackets", "))", false},
		{"edge case: only opening brackets", "((", false},
		{"edge case: closer before opener", ")(", false},
		{"edge case: deep nesting", "{[()]}", true},
		{"edge case: unclosed opener after valid prefix", "()((", false},
		{"edge case: mixed nesting and sequence", "{[]}([]){}", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isValid(tt.input)
			if result != tt.expected {
				t.Errorf("isValid(%q) = %v, want %v", tt.input, result, tt.expected)
			}
		})
	}
}
