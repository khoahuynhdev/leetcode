package main

import "testing"

func TestSolution(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		expected int
	}{
		{"example 1: second open fully closed, first needs one more", "(()))", 1},
		{"example 2: already balanced", "())", 0},
		{"example 3: leading closers and trailing opener", "))())(", 3},
		{"edge case: single open", "(", 2},
		{"edge case: single close", ")", 2},
		{"edge case: only openers", "(((", 6},
		{"edge case: odd closers with no openers", ")))", 3},
		{"edge case: lone closers separated by openers", "()()", 2},
		{"edge case: nested balanced", "(())))", 0},
		{"edge case: balanced long example", "())(())))", 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := minInsertions(tt.s)
			if result != tt.expected {
				t.Errorf("minInsertions(%q) = %d, want %d", tt.s, result, tt.expected)
			}
		})
	}
}
